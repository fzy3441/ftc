package tt

func Ternary(check bool,left,right interface{}) interface{} {
	if check {
		return left
	}else
	{
		return right
	}
}
