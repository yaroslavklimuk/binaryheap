package binaryheap

import "testing"

func TestBinaryHeap_Insert(t *testing.T) {
	type fields struct {
		internalSlice []int
		swap          int
		maxTop        bool
	}
	tests := []struct {
		name      string
		fields    fields
		toInsert  []int
		expResult []int
	}{
		{
			name: "first case",
			fields: fields{
				internalSlice: []int{42, 29, 18, 14, 7, 18, 12, 11, 13},
				maxTop:        true,
			},
			toInsert:  []int{35},
			expResult: []int{42, 35, 18, 14, 29, 18, 12, 11, 13, 7},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bh := &BinaryHeap{
				internalSlice: tt.fields.internalSlice,
				swap:          tt.fields.swap,
				maxTop:        tt.fields.maxTop,
			}
			for _, item := range tt.toInsert {
				bh.Insert(item)
			}
			if len(bh.internalSlice) != len(tt.expResult) {
				t.Errorf("The resulting slice has an unexpected length."+
					"Expected: %d, actual: %d",
					len(tt.expResult), len(bh.internalSlice))
				return
			}
			for i := 0; i < len(bh.internalSlice); i++ {
				if bh.internalSlice[i] != tt.expResult[i] {
					t.Errorf("The resulting slice has an unexpected content."+
						"Index: %d, expected: %d, got: %d",
						i, tt.expResult[i], bh.internalSlice[i])
				}
			}
		})
	}
}

func TestBinaryHeap_PopTop(t *testing.T) {
	type fields struct {
		internalSlice []int
		swap          int
		maxTop        bool
	}
	tests := []struct {
		name      string
		fields    fields
		want      int
		expResult []int
	}{
		{
			name: "first case",
			fields: fields{
				internalSlice: []int{42, 35, 18, 14, 29, 18, 12, 11, 13, 7},
				maxTop:        true,
			},
			want:      42,
			expResult: []int{35, 29, 18, 14, 7, 18, 12, 11, 13},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bh := &BinaryHeap{
				internalSlice: tt.fields.internalSlice,
				swap:          tt.fields.swap,
				maxTop:        tt.fields.maxTop,
			}
			if got := bh.PopTop(); got != tt.want {
				t.Errorf("PopTop() = %v, want %v", got, tt.want)
			}
			if len(bh.internalSlice) != len(tt.expResult) {
				t.Errorf("The resulting slice has an unexpected length."+
					"Expected: %d, actual: %d",
					len(tt.expResult), len(bh.internalSlice))
				return
			}
			for i := 0; i < len(bh.internalSlice); i++ {
				if bh.internalSlice[i] != tt.expResult[i] {
					t.Errorf("The resulting slice has an unexpected content."+
						"Index: %d, expected: %d, got: %d",
						i, tt.expResult[i], bh.internalSlice[i])
				}
			}
		})
	}
}
