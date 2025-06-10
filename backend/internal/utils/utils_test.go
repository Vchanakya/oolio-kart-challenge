package utils

import "testing"

func TestHasCouponInAtLeastTwo(t *testing.T) {
	testFiles := []string{
		"couponbase1.gz",
		"couponbase2.gz",
		"couponbase3.gz",
	}
	if !HasCouponInAtLeastTwo(testFiles, "HAPPYHRS") {
		t.Fatal("expected true, got false")
	}
	if HasCouponInAtLeastTwo(testFiles, "SUPER100") {
		t.Fatal("expected false, got true")
	}
}
