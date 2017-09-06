package glplus

import (
	"testing"
)

func TestLayout(t *testing.T) {
	// Create our root item. Items are just 2D boxes.
	ctx := Lay_context{}

	// Create our root item. Items are just 2D boxes.
	root := ctx.Lay_item()

	// Let's pretend we have a window in our game or OS of some known dimension.
	// We'll want to explicitly set our root item to be that size.
	ctx.Lay_set_size_xy(root, 1280, 720)

	// Set our root item to arrange its children in a row, left-to-right, in the
	// order they are inserted.
	ctx.Lay_set_contain(root, LAY_ROW)

	// Create the item for our master list.
	master_list := ctx.Lay_item()
	ctx.Lay_insert(root, master_list)

	// Our master list has a specific fixed width, but we want it to fill all
	// available vertical space.
	ctx.Lay_set_size_xy(master_list, 400, 0)
	// We set our item's behavior within its parent to desire filling up available
	// vertical space.
	ctx.Lay_set_behave(master_list, LAY_VFILL)
	// And we set it so that it will lay out its children in a column,
	// top-to-bottom, in the order they are inserted.
	ctx.Lay_set_contain(master_list, LAY_COLUMN)

	content_view := ctx.Lay_item()
	ctx.Lay_insert(root, content_view)
	// The content view just wants to fill up all of the remaining space, so we
	// don't need to set any size on it.
	//
	// We could just set LAY_FILL here instead of bitwise-or'ing LAY_HFILL and
	// LAY_VFILL, but I want to demonstrate that this is how you combine flags.
	ctx.Lay_set_behave(content_view, LAY_HFILL|LAY_VFILL)

	// Normally at this point, we would probably want to create items for our
	// master list and our content view and insert them. This is just a dumb fake
	// example, so let's move on to finishing up.

	// Run the context -- this does all of the actual calculations.
	ctx.Lay_run_context()

	// Now we can get the calculated size of our items as 2D rectangles. The four
	// components of the vector represent x and y of the top left corner, and then
	// the width and height.
	master_list_rect := ctx.Lay_get_rect(master_list)
	content_view_rect := ctx.Lay_get_rect(content_view)

	if (master_list_rect != Lay_vec4{0, 0, 400, 720}) {
		t.Fail()
	}
	if (content_view_rect != Lay_vec4{400, 0, 880, 720}) {
		t.Fail()
	}
}

func LTEST_TRUE(t *testing.T, cond bool) {
	if !cond {
		t.Fail()
	}
}

func TestSimpleFill(t *testing.T) {
	ctx := Lay_context{}
	root := ctx.Lay_item()
	child := ctx.Lay_item()

	ctx.Lay_set_size_xy(root, 30, 40)
	ctx.Lay_set_behave(child, LAY_FILL)
	ctx.Lay_insert(root, child)

	ctx.Lay_run_context()

	root_r := ctx.Lay_get_rect(root)
	child_r := ctx.Lay_get_rect(child)

	LTEST_TRUE(t, root_r[0] == 0 && root_r[1] == 0)
	LTEST_TRUE(t, root_r[2] == 30 && root_r[3] == 40)

	LTEST_TRUE(t, child_r[0] == 0 && child_r[1] == 0)
	LTEST_TRUE(t, child_r[2] == 30 && child_r[3] == 40)

	// Test to make sure size is ok
	root_size := ctx.Lay_get_size(root)
	// Make sure _xy version gives the same values
	root_size_cs := ctx.Lay_get_size_xy(root)
	LTEST_TRUE(t, root_size[0] == 30 && root_size[1] == 40)
	LTEST_TRUE(t, root_size[0] == root_size_cs[0] && root_size[1] == root_size_cs[1])

	// Test to make sure the _xywh getter produces the same results as the
	// ctx.Lay_vec4-return-value version.

	root_r_cs := ctx.Lay_get_rect_xywh(root)
	child_r_cs := ctx.Lay_get_rect_xywh(child)
	LTEST_TRUE(t, root_r[0] == root_r_cs[0] && root_r[1] == root_r_cs[1] && root_r[2] == root_r_cs[2] && root_r[3] == root_r_cs[3])
	LTEST_TRUE(t, child_r[0] == child_r_cs[0] && child_r[1] == child_r_cs[1] && child_r[2] == child_r_cs[2] && child_r[3] == child_r_cs[3])
}
