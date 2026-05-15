package dataset

func coarseKey(v [14]float32) uint32 {

	a := uint32(v[0] * 10)
	b := uint32(v[2] * 10)
	c := uint32(v[7] * 10)
	d := uint32(v[12] * 10)

	return (a << 24) | (b << 16) | (c << 8) | d

	// amountBucket := int(v[0] * 20)
	// hourBucket := int(v[3] * 6)
	// mccBucket := int(v[12] * 10)

	// flags :=
	// 	int(v[9])*4 +
	// 	int(v[10])*2 +
	// 	int(v[11])

	// return fmt.Sprintf(
	// 	"%02d|%02d|%02d|%d",
	// 	amountBucket,
	// 	hourBucket,
	// 	mccBucket,
	// 	flags,
	// )
}

func broadKey(v [14]float32) uint32 {

	a := uint32(v[0] * 5)
	b := uint32(v[2] * 5)

	return (a << 8) | b

	// mccBucket := int(v[12] * 10)

	// flags :=
	// 	int(v[9])*4 +
	// 	int(v[10])*2 +
	// 	int(v[11])

	// return fmt.Sprintf(
	// 	"%02d|%d",
	// 	mccBucket,
	// 	flags,
	// )
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