package disjointset_test

import (
	"testing"

	ds "dsa/ds/disjoint-set" // Corrected module path

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

func TestMakeSet(t *testing.T) {
	dsObj := ds.NewDisjointSet() 
	dsObj.MakeSet(1)

	if _, exists := dsObj.Elements[1]; !exists { 
		t.Fatalf("MakeSet(1) failed, element not found in map")
	}

	got := dsObj.Elements[1] 
	// Create expected element, Parent will be checked separately
	want := &ds.Element{Value: 1, Rank: 0} 

	// Allow comparing unexported Parent field by ignoring it initially
	// We need to compare pointers for Parent check later.
	opts := cmpopts.IgnoreFields(ds.Element{}, "Parent") 

	if diff := cmp.Diff(want, got, opts); diff != "" {
		t.Errorf("MakeSet(1) returned diff (-want +got):\n%s", diff)
	}
	// Check that the parent points to itself
	if got.Parent != got {
		t.Errorf("MakeSet(1): Parent: got %p, want %p (itself)", got.Parent, got)
	}

	// Test making a set that already exists
	initialElementPtr := dsObj.Elements[1] 
	dsObj.MakeSet(1) 
	if dsObj.Elements[1] != initialElementPtr { 
		t.Errorf("MakeSet(1) on existing element created a new element instance")
	}
	if len(dsObj.Elements) != 1 { 
		 t.Errorf("MakeSet(1) on existing element changed map size: got %d, want 1", len(dsObj.Elements))
	}
}

func TestFindSet(t *testing.T) {
	dsObj := ds.NewDisjointSet() 

	// Test finding non-existent element
	if got := dsObj.FindSet(1); got != nil {
		t.Errorf("FindSet(1) on empty set: got %v, want nil", got)
	}

	dsObj.MakeSet(1)
	dsObj.MakeSet(2)
	dsObj.MakeSet(3)

	e1 := dsObj.Elements[1] 
	e2 := dsObj.Elements[2] 
	e3 := dsObj.Elements[3] 

	// Create a chain: 1 -> 2 -> 3 (root)
	e1.Parent = e2
	e2.Parent = e3

	// Find element 1, should return root 3
	gotRoot := dsObj.FindSet(1)
	if gotRoot != e3 {
		t.Errorf("FindSet(1): got root %v, want root %v", gotRoot, e3)
	}

	// Check path compression: both 1 and 2 should now point directly to 3
	if e1.Parent != e3 {
		t.Errorf("Path compression failed for element 1: Parent: got %v, want %v", e1.Parent, e3)
	}
	if e2.Parent != e3 {
		t.Errorf("Path compression failed for element 2: Parent: got %v, want %v", e2.Parent, e3)
	}

	// Find root element 3 (should return itself)
	gotRoot = dsObj.FindSet(3)
	if gotRoot != e3 {
		t.Errorf("FindSet(3): got root %v, want root %v", gotRoot, e3)
	}
	if e3.Parent != e3 {
		t.Errorf("FindSet(3): Root's parent should be itself: got %v, want %v", e3.Parent, e3)
	}
}


func TestUnion(t *testing.T) {
	dsObj := ds.NewDisjointSet() 

	dsObj.MakeSet(1)
	dsObj.MakeSet(2)
	dsObj.MakeSet(3)
	dsObj.MakeSet(4)
	dsObj.MakeSet(5)

	e1 := dsObj.Elements[1] 
	e2 := dsObj.Elements[2] 
	e3 := dsObj.Elements[3]
	e4 := dsObj.Elements[4]
	// e5 is unused but let's get it for consistency if needed later
	// e5 := dsObj.Elements[5] 


	// --- Union 1 and 2 (rank 0 + rank 0) ---
	dsObj.Union(1, 2)
	root1_2 := dsObj.FindSet(1) // performs path compression
	wantRank1_2 := 1
	if root1_2 != dsObj.FindSet(2) {
		t.Fatalf("Union(1, 2) failed, elements not in the same set")
	}
	if root1_2.Rank != wantRank1_2 {
		t.Errorf("Union(1, 2): Rank: got %d, want %d", root1_2.Rank, wantRank1_2)
	}
	// Check which element became the parent (depends on implementation detail, either is fine)
	// Re-find parents after potential path compression during FindSet(2)
	if !(e1.Parent == root1_2 && e2.Parent == root1_2) {
		 t.Errorf("Union(1, 2): Incorrect parent assignment after FindSet")
	}


	// --- Union 3 and 4 (rank 0 + rank 0) ---
	dsObj.Union(3, 4)
	root3_4 := dsObj.FindSet(3) // performs path compression
	wantRank3_4 := 1
	if root3_4 != dsObj.FindSet(4) {
		t.Fatalf("Union(3, 4) failed, elements not in the same set")
	}
	if root3_4.Rank != wantRank3_4 {
		t.Errorf("Union(3, 4): Rank: got %d, want %d", root3_4.Rank, wantRank3_4)
	}
	// Re-find parents after potential path compression during FindSet(4)
	if !(e3.Parent == root3_4 && e4.Parent == root3_4) {
		 t.Errorf("Union(3, 4): Incorrect parent assignment after FindSet")
	}


	// --- Union sets {1, 2} and {3, 4} (rank 1 + rank 1) ---
	dsObj.Union(1, 3)
	finalRoot := dsObj.FindSet(1) // Find one element to get the final root
	wantFinalRank := 2
	if finalRoot.Rank != wantFinalRank {
		 t.Errorf("Union(1, 3): Final Rank: got %d, want %d", finalRoot.Rank, wantFinalRank)
	}
	// Check all elements now point to the same root
	for _, val := range []int{1, 2, 3, 4} {
		if gotRoot := dsObj.FindSet(val); gotRoot != finalRoot {
			t.Errorf("Union(1, 3): Element %d not pointing to final root %v, got %v", val, finalRoot, gotRoot)
		}
	}


	// --- Union with itself (should be no-op) ---
	rootBeforeSelfUnion := dsObj.FindSet(1)
	rankBeforeSelfUnion := rootBeforeSelfUnion.Rank
	dsObj.Union(1, 1)
	rootAfterSelfUnion := dsObj.FindSet(1)
	if rootAfterSelfUnion != rootBeforeSelfUnion {
		t.Errorf("Union(1, 1) changed root: got %v, want %v", rootAfterSelfUnion, rootBeforeSelfUnion)
	}
	if rootAfterSelfUnion.Rank != rankBeforeSelfUnion {
		t.Errorf("Union(1, 1) changed rank: got %d, want %d", rootAfterSelfUnion.Rank, rankBeforeSelfUnion)
	}


	// --- Union with non-existent element (should be no-op) ---
	rootBeforeNonExistentUnion := dsObj.FindSet(1)
	rankBeforeNonExistentUnion := rootBeforeNonExistentUnion.Rank
	dsObj.Union(1, 6) // 6 doesn't exist
	rootAfterNonExistentUnion := dsObj.FindSet(1)
	if rootAfterNonExistentUnion != rootBeforeNonExistentUnion {
		t.Errorf("Union(1, 6) changed root: got %v, want %v", rootAfterNonExistentUnion, rootBeforeNonExistentUnion)
	}
	if rootAfterNonExistentUnion.Rank != rankBeforeNonExistentUnion {
		t.Errorf("Union(1, 6) changed rank: got %d, want %d", rootAfterNonExistentUnion.Rank, rankBeforeNonExistentUnion)
	}
	dsObj.Union(6, 1) // Try the other way around
	rootAfterNonExistentUnion2 := dsObj.FindSet(1)
	if rootAfterNonExistentUnion2 != rootBeforeNonExistentUnion {
		t.Errorf("Union(6, 1) changed root: got %v, want %v", rootAfterNonExistentUnion2, rootBeforeNonExistentUnion)
	}
	if rootAfterNonExistentUnion2.Rank != rankBeforeNonExistentUnion {
		t.Errorf("Union(6, 1) changed rank: got %d, want %d", rootAfterNonExistentUnion2.Rank, rankBeforeNonExistentUnion)
	}


	// --- Test union by rank specifically ---
	dsObj.MakeSet(6)
	dsObj.MakeSet(7) // unused
	e6 := dsObj.Elements[6] 
	e6.Rank = 3 // Make rank of 6 higher than {1,2,3,4}'s root (rank 2)

	root1BeforeRankUnion := dsObj.FindSet(1)

	dsObj.Union(1, 6) // Union set {1,2,3,4} (rank 2) with {6} (rank 3)
	newRoot := dsObj.FindSet(1) // Should now be e6
	wantNewRoot := e6
	wantNewRank := 3 // Rank shouldn't increase

	if newRoot != wantNewRoot {
		 t.Errorf("Union by rank (1, 6): Root: got %v, want %v", newRoot, wantNewRoot)
	}
	if newRoot.Rank != wantNewRank {
		t.Errorf("Union by rank (1, 6): Rank: got %d, want %d", newRoot.Rank, wantNewRank)
	}
	if root1BeforeRankUnion.Parent != newRoot { // Check old root points to new root
		 t.Errorf("Union by rank (1, 6): Old root's parent: got %v, want %v", root1BeforeRankUnion.Parent, newRoot)
	}

	// --- Test union of equal, non-zero ranks ---
	dsObj.MakeSet(8)
	dsObj.MakeSet(9)
	e8 := dsObj.Elements[8] 
	e9 := dsObj.Elements[9] 
	e8.Rank = 5
	e9.Rank = 5

	dsObj.Union(8, 9)
	root8_9 := dsObj.FindSet(8)
	wantRank8_9 := 6 // Rank should increase

	if root8_9 != dsObj.FindSet(9) {
		t.Fatalf("Union(8, 9) of equal ranks failed, elements not in the same set")
	}
	if root8_9.Rank != wantRank8_9 {
		 t.Errorf("Union(8, 9) of equal ranks: Rank: got %d, want %d", root8_9.Rank, wantRank8_9)
	}
	// Check parent assignment (either e8 or e9 becomes the root)
	// Re-find parents after potential path compression during FindSet(9)
	if !(e8.Parent == root8_9 && e9.Parent == root8_9) {
		t.Errorf("Union(8, 9): Incorrect parent assignment after FindSet")
	}
}
