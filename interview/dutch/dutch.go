package dutch

type dutch struct {
	p1Idx, p2Idx, p3Idx       int
	p1Color, p2Color, p3Color string
}

// Sort the buckets by color
// Assumes the buckets array only contains 0's, 1's or 2's
// which represents the Dutch flag colors
func Sort(buckets []int) {
	// create dutch struct and define partition indexes
	flag := dutch{
		p1Idx:   1,
		p2Idx:   len(buckets) / 3,
		p3Idx:   len(buckets) / 3 * 2,
		p1Color: color(0),
	}

	flag.init(buckets)
	for i := flag.p1Idx; i < len(buckets)/3; i++ {
		// get destination bucket
		// swap with current idx at destination bucket
	}
}

func (flag *dutch) init(buckets []int) {
	if p2IdxColor := color(buckets[flag.p2Idx]); flag.p1Color == p2IdxColor {
		flag.putInDestinationBucket(buckets, flag.p2Idx)
	} else {
		flag.p2Color = p2IdxColor
	}

	if p3IdxColor := color(buckets[flag.p3Idx]); flag.p1Color == p3IdxColor {
		flag.putInDestinationBucket(buckets, flag.p3Idx)
	} else if flag.p2Color == p3IdxColor {
		flag.putInDestinationBucket(buckets, flag.p3Idx)
	} else {
		flag.p3Color = p3IdxColor
	}
}

func (flag *dutch) putInDestinationBucket(buckets []int, i int) {
	// get destination bucket
	dest := flag.getDestinationBucket(color(i))
	if dest != -1 {
		swap(buckets, i, dest)
	}
}

func (flag *dutch) getDestinationBucket(color string) int {
	switch color {
	case flag.p1Color:
		return flag.p1Idx
	case flag.p2Color:
		return flag.p2Idx
	case flag.p3Color:
		return flag.p3Idx
	default:
		return -1
	}
}

// returns the color of the bucket
func color(i int) string {
	switch i {
	case 0: // `b`
		return "blue"
	case 1: // `w`
		return "white"
	case 2: // `r`
		return "red"
	default:
		panic("unrecognized color in bucket")
	}
}

func swap(array []int, i, j int) {
	array[i], array[j] = array[j], array[i]
}
