package nom

func Delimited[T any](start AnyParser, value Parser[T], end AnyParser) Parser[T] {
	return func(input string) (string, T, error) {
		var item T
		input, _, err := start(input)
		if err != nil {
			return input, item, err
		}
		input, item, err = value(input)
		if err != nil {
			return input, item, err
		}
		input, _, err = end(input)
		return input, item, err
	}
}

func Preceded[T any](start AnyParser, value Parser[T]) Parser[T] {
	return func(input string) (string, T, error) {
		input, _, err := start(input)
		if err != nil {
			var item T
			return input, item, err
		}
		input, item, err := value(input)
		return input, item, err
	}
}

func Terminated[T any](value Parser[T], end AnyParser) Parser[T] {
	return func(input string) (string, T, error) {
		input, item, err := value(input)
		if err != nil {
			return input, item, err
		}
		input, _, err = end(input)
		return input, item, err
	}
}

type Tuple2[T1 any, T2 any] struct {
	First T1
	Second T2
}

func Pair2[T1 any, T2 any](first Parser[T1], second Parser[T2]) Parser[Tuple2[T1, T2]] {
	return func(input string) (string, Tuple2[T1, T2], error) {
		input, firstItem, err := first(input)
		if err != nil {
			return input, Tuple2[T1, T2]{First: firstItem}, err
		}
		input, secondItem, err := second(input)
		return input, Tuple2[T1, T2]{First: firstItem, Second: secondItem}, err
	}
}

func SeparatedPair2[T1 any, T2 any](first Parser[T1], sep AnyParser, second Parser[T2]) Parser[Tuple2[T1, T2]] {
	return func(input string) (string, Tuple2[T1, T2], error) {
		input, firstItem, err := first(input)
		if err != nil {
			return input, Tuple2[T1, T2]{First: firstItem}, err
		}
		input, _, err = sep(input)
		if err != nil {
			return input, Tuple2[T1, T2]{First: firstItem}, err
		}
		input, secondItem, err := second(input)
		return input, Tuple2[T1, T2]{First: firstItem, Second: secondItem}, err
	}
}
