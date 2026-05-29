package dataset

//import "fmt"

const maxCandidates = 128

func limitCandidates(ids []uint32, max int) []uint32 {
	if len(ids) <= max {
		return ids
	}

	out := make([]uint32, 0, max)
	step := len(ids) / max
	if step < 1 {
		step = 1
	}

	for i := 0; i < len(ids) && len(out) < max; i += step {
		out = append(out, ids[i])
	}

	return out
}

func coarseKey(v [14]float32) uint32 {

	amount := uint32(v[0] * 8)   // more granular
	hour := uint32(v[3] * 6)
	day := uint32(v[4] * 7)
	tx := uint32(v[8] * 4)
	mcc := uint32(v[12] * 10)

	flags := uint32(v[9])<<2 | uint32(v[10])<<1 | uint32(v[11])

	return (amount << 24) | (hour << 20) | (day << 16) | (tx << 12) | (flags << 8) | mcc


	// a := uint32(v[0] * 10)
	// b := uint32(v[2] * 10)
	// c := uint32(v[7] * 10)
	// d := uint32(v[12] * 10)

	// return (a << 24) | (b << 16) | (c << 8) | d

}

func broadKey(v [14]float32) uint32 {
	amount := uint32(v[0] * 4)
	tx := uint32(v[8] * 2)
	flags := uint32(v[9])<<2 | uint32(v[10])<<1 | uint32(v[11])

	return (amount << 8) | (tx << 4) | flags

	// less buckets, broader buckets
	// a := uint32(v[0] * 2)
	// b := uint32(v[2] * 2)

	// return (a << 8) | b

}

func (d *Dataset) Candidates(q [14]float32) []uint32 {

	ck := coarseKey(q)

	if ids := d.CoarseIndex[ck]; len(ids) > 0 {
		return limitCandidates(ids, maxCandidates)
	}

	bk := broadKey(q)

	if ids := d.BroadIndex[bk]; len(ids) > 0 {
		return limitCandidates(ids, maxCandidates)
	}

	// if nil, any qyery without vectors could try to 3M search --> no full scan
	// return d.BroadIndex[broadKey(q)] 
	return nil
}