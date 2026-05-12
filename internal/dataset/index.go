package dataset

import "fmt"

func coarseKey(v [14]float32) string {

	amountBucket := int(v[0] * 20)
	hourBucket := int(v[3] * 6)
	mccBucket := int(v[12] * 10)

	flags :=
		int(v[9])*4 +
		int(v[10])*2 +
		int(v[11])

	return fmt.Sprintf(
		"%02d|%02d|%02d|%d",
		amountBucket,
		hourBucket,
		mccBucket,
		flags,
	)
}

func broadKey(v [14]float32) string {

	mccBucket := int(v[12] * 10)

	flags :=
		int(v[9])*4 +
		int(v[10])*2 +
		int(v[11])

	return fmt.Sprintf(
		"%02d|%d",
		mccBucket,
		flags,
	)
}

func (d *Dataset) Candidates(q [14]float32) []int {

	ck := coarseKey(q)

	if ids := d.CoarseIndex[ck]; len(ids) >= 32 {
		return ids
	}

	bk := broadKey(q)

	if ids := d.BroadIndex[bk]; len(ids) > 0 {
		return ids
	}

	return nil
}