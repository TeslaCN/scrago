package dedup

/*
Java 8 String.hashCode()
public int hashCode() {
	int h = hash;
	if (h == 0 && value.length > 0) {
		char val[] = value;

		for (int i = 0; i < value.length; i++) {
			h = 31 * h + val[i];
		}
		hash = h;
	}
	return h;
}
*/

func HashCode(s string) int32 {
	var h int32 = 0
	if len(s) > 0 {
		for _, c := range s {
			h = 31*h + c
		}
	}
	return h
}

func Reserve(i int32, length int32) int32 {
	return ^(^0 << length) & i
}
