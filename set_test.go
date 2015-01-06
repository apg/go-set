package set

import (
	"testing"
)

func TestBasicSet(t *testing.T) {
	ut := New()
	if ut.Len() != 0 {
		t.Errorf("New should return an empty set.")
	}

	if !ut.Add(1) {
		t.Errorf("Add to an empty set return false.")
	}

	if ut.Len() != 1 {
		t.Errorf("Add did not increase cardinality of set by 1")
	}

	if !ut.Member(1) {
		t.Errorf("Member did not return true, but element was added.")
	}

	if !ut.Delete(1) {
		t.Errorf("Delete did not return true")
	}

	if ut.Len() != 0 {
		t.Errorf("Delete did not decrease cardinality of set by 1")
	}

	if ut.Member(1) {
		t.Errorf("Member did not returned true, but element should have been deleted.")
	}

}

func TestCopySet(t *testing.T) {
	ut := New(1, 2, 3)
	ut2 := ut.Copy()

	if ut2.Len() == 0 {
		t.Errorf("Copy should produce a non empty set.")
	}

	if ut.Len() != ut2.Len() {
		t.Errorf("Copy should result in a set with equal cardinality.")
	}

	for _, x := range ut.Freeze() {
		if !ut2.Member(x) {
			t.Errorf("Item from original was not found in copy.")
		}
	}
}

func TestUpdateSet(t *testing.T) {
	ut := New(1, 2, 3)
	ut.Merge(New(2, 3, 4))

	if ut.Len() != 4 {
		t.Errorf("Merge of an overlapping set should add only non-overlapping elements")
	}

	ut.Discard(ut.Copy())
	if !ut.Empty() {
		t.Errorf("Discard of a copy of the set should produce the empty set")
	}
}

func TestNondestructiveUpdateSet(t *testing.T) {
	ut := New(1, 2, 3)
	ut2 := New(2, 3, 4)

	tmp := ut.Union(ut2)

	if tmp.Empty() {
		t.Errorf("Union of two non-empty sets is not empty.")
	}

	if tmp.Len() != 4 {
		t.Errorf("Union of {1, 2, 3} and {2, 3, 4} should have a cardinality of 4")
	}

	for _, k := range ut2.Freeze() {
		if !tmp.Member(k) {
			t.Errorf("Expected %u to be a member of union.", k)
		}
	}

	tmp2 := ut.Difference(ut2)

	if tmp2.Len() != 1 {
		t.Errorf("Expected difference to be just 1 element.")
	}

	if !tmp2.Member(1) {
		t.Errorf("Expected 1 to be a member of difference.")
	}

	tmp3 := ut.Intersection(ut2)
	if tmp3.Len() != 2 {
		t.Errorf("Expected intersection to be 2 elements.")
	}

	if !tmp3.Member(2) || !tmp3.Member(3) {
		t.Errorf("Expected both 2 and 3 to be members of intersection.")
	}
}

func TestSubSuperSet(t *testing.T) {
	ut := New(1, 2, 3)
	ut2 := New(1, 2, 3, 4)
	if !ut2.Superset(ut) {
		t.Errorf("{1, 2, 3, 4} is a superset of {1, 2, 3}")
	}

	if !ut.Subset(ut2) {
		t.Errorf("{1, 2, 3} is a subset of {1, 2, 3, 4}")
	}

}
