// TestSum for testing
func TestSum(t *testing.T) {

	fn := GenDisplaceFn(10, 2, 1)
	//fmt.Println(fn(3))
	//fmt.Println(fn(5))

	total := (fn(3))
	if total != 20.5 {
		t.Errorf("Sum was incorrect, got: %f, want: %f.", total, 20.5)
	}

	total = (fn(5))
	if total != 32.5 {
		t.Errorf("Sum was incorrect, got: %f, want: %f.", total, 32.5)
	}
}