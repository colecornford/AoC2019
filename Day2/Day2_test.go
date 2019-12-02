package day2

import "testing"

func TestRestoreState(t *testing.T) {
	
	// todo: DRY here + add loop.
	// todo: test data into a [][]int

	testSuite := [][]int{{1,0,0,0,99},{2,3,0,3,99},{2,4,4,5,99,0},{1,1,1,4,99,5,6,0,99}}
	expectedSuite := [][]int{{2,0,0,0,99},{2,3,0,6,99},{2,4,4,5,99,9801},{30,1,1,4,2,5,6,0,99}}

	for index, test := range testSuite {
		result := restoreState(test)
		for subIndex, num := range result {
			if num != expectedSuite[index][subIndex] {
				t.Errorf("Failed on: %v. Got: %v. Expected: %v", test, result, expectedSuite[index])
			}
		}
	}
}
