package tt

func Ternary(check bool,left interface{},right interface{}) interface{} {
	if check {
		return left
	}else
	{
		return right
	}
}
