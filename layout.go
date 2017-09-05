package glplus

// https://github.com/randrew/layout

import (
	"fmt"
	"math"
)

type lay_id uint32
type lay_scalar int16

const LAY_INVALID_ID lay_id = math.MaxUint32

type lay_vec4 [4]lay_scalar
type lay_vec2 [2]lay_scalar

type lay_item_t struct {
	flags        int32
	first_child  lay_id
	next_sibling lay_id
	margins      lay_vec4
	size         lay_vec2
}

type lay_context struct {
	items []*lay_item_t
	rects []lay_vec4
}

// Container flags to pass to lay_set_container()
const (
	// flex-direction (bit 0+1)

	// left to right
	LAY_ROW = 0x002
	// top to bottom
	LAY_COLUMN = 0x003

	// model (bit 1)

	// free layout
	LAY_LAYOUT = 0x000
	// flex model
	LAY_FLEX = 0x002

	// flex-wrap (bit 2)

	// single-line
	LAY_NOWRAP = 0x000
	// multi-line, wrap left to right
	LAY_WRAP = 0x004

	// justify-content (start, end, center, space-between)
	// at start of row/column
	LAY_START = 0x008
	// at center of row/column
	LAY_MIDDLE = 0x000
	// at end of row/column
	LAY_END = 0x010
	// insert spacing to stretch across whole row/column
	LAY_JUSTIFY = 0x018

	// align-items
	// can be implemented by putting a flex container in a layout container,
	// then using LAY_TOP, LAY_BOTTOM, LAY_VFILL, LAY_VCENTER, etc.
	// FILL is equivalent to stretch/grow

	// align-content (start, end, center, stretch)
	// can be implemented by putting a flex container in a layout container,
	// then using LAY_TOP, LAY_BOTTOM, LAY_VFILL, LAY_VCENTER, etc.
	// FILL is equivalent to stretch space-between is not supported.
)

type lay_box_flags int

// child layout flags to pass to lay_set_behave()
const (
	// attachments (bit 5-8)
	// fully valid when parent uses LAY_LAYOUT model
	// partially valid when in LAY_FLEX model

	// anchor to left item or left side of parent
	LAY_LEFT = 0x020
	// anchor to top item or top side of parent
	LAY_TOP = 0x040
	// anchor to right item or right side of parent
	LAY_RIGHT = 0x080
	// anchor to bottom item or bottom side of parent
	LAY_BOTTOM = 0x100
	// anchor to both left and right item or parent borders
	LAY_HFILL = 0x0a0
	// anchor to both top and bottom item or parent borders
	LAY_VFILL = 0x140
	// center horizontally, with left margin as offset
	LAY_HCENTER = 0x000
	// center vertically, with top margin as offset
	LAY_VCENTER = 0x000
	// center in both directions, with left/top margin as offset
	LAY_CENTER = 0x000
	// anchor to all four directions
	LAY_FILL = 0x1e0
	// When in a wrapping container, put this element on a new line. Wrapping
	// layout code auto-inserts LAY_BREAK flags as needed. See GitHub issues for
	// TODO related to this.
	//
	// Drawing routines can read this via item pointers as needed after
	// performing layout calculations.
	LAY_BREAK = 0x200
)

type lay_layout_flags int

const (
	// these bits, starting at bit 16, can be safely assigned by the
	// application, e.g. as item types, other event types, drop targets, etc.
	// this is not yet exposed via API functions, you'll need to get/set these
	// by directly accessing item pointers.
	//
	// (In reality we have more free bits than this, TODO)
	//
	// TODO fix int/unsigned size mismatch (clang issues warning for this),
	// should be all bits as 1 instead of INT_MAX
	LAY_USERMASK = 0x7fff0000

	// a special mask passed to lay_find_item() (currently does not exist, was
	// not ported from oui)
	LAY_ANY = 0x7fffffff
)

const (
	// extra item flags

	// bit 0-2
	LAY_ITEM_BOX_MODEL_MASK = 0x000007
	// bit 0-4
	LAY_ITEM_BOX_MASK = 0x00001F
	// bit 5-9
	LAY_ITEM_LAYOUT_MASK = 0x0003E0
	// item has been inserted (bit 10)
	LAY_ITEM_INSERTED = 0x400
	// horizontal size has been explicitly set (bit 11)
	LAY_ITEM_HFIXED = 0x800
	// vertical size has been explicitly set (bit 12)
	LAY_ITEM_VFIXED = 0x1000
	// bit 11-12
	LAY_ITEM_FIXED_MASK = LAY_ITEM_HFIXED | LAY_ITEM_VFIXED

	// which flag bits will be compared
	LAY_ITEM_COMPARE_MASK = LAY_ITEM_BOX_MODEL_MASK | (LAY_ITEM_LAYOUT_MASK & ^LAY_BREAK) | LAY_USERMASK
)

func lay_vec4_xyzw(x, y, z, w lay_scalar) lay_vec4 {
	return lay_vec4{x, y, z, w}
}

func LAY_ASSERT(cond bool) {
	if !cond {
		panic(fmt.Errorf("LAY_ASSERT\n"))
	}
}

// Get the pointer to an item in the buffer by its id. Don't keep this around --
// it will become invalid as soon as any reallocation occurs. Just store the id
// instead (it's smaller, anyway, and the lookup cost will be nothing.)
func (ctx *lay_context) lay_get_item(id lay_id) *lay_item_t {
	LAY_ASSERT(id != LAY_INVALID_ID && id < lay_id(len(ctx.items)))
	return ctx.items[id]
}

// Get the id of first child of an item, if any. Returns LAY_INVALID_ID if there
// is no child.
func (ctx *lay_context) lay_first_child(id lay_id) lay_id {
	pitem := ctx.lay_get_item(id)
	return pitem.first_child
}

// Get the id of the next sibling of an item, if any. Returns LAY_INVALID_ID if
// there is no next sibling.
func (ctx *lay_context) lay_next_sibling(id lay_id) lay_id {
	pitem := ctx.lay_get_item(id)
	return pitem.next_sibling
}

// Returns the calculated rectangle of an item. This is only valid after calling
// lay_run_context and before any other reallocation occurs. Otherwise, the
// result will be undefined. The vector components are:
// 0: x starting position, 1: y starting position
// 2: width, 3: height
func (ctx *lay_context) lay_get_rect(id lay_id) lay_vec4 {
	LAY_ASSERT(id != LAY_INVALID_ID && id < lay_id(len(ctx.items)))
	return ctx.rects[id]
}

// The same as lay_get_rect, but writes the x,y positions and width,height
// values to the specified addresses instead of returning them in a lay_vec4.
func (ctx *lay_context) lay_get_rect_xywh(id lay_id) lay_vec4 {
	LAY_ASSERT(id != LAY_INVALID_ID && id < lay_id(len(ctx.items)))
	return ctx.rects[id]
}

// Notes about the use of vector_size merely for syntax convenience:
//
// The current layout calculation procedures are not written in a way that
// would benefit from SIMD instruction usage.
//
// (Passing 128-bit float4 vectors using __vectorcall *might* get you some
// small benefit in very specific situations, but is unlikely to be worth the
// hassle. And I believe this would only be needed if you compiled the library
// in a way where the compiler was prevented from using inlining when copying
// rectangle/size data.)
//
// I might go back in the future and just use regular struct-wrapped arrays.
// I'm not sure if relying the vector thing in GCC/clang and then using C++
// operator overloading in MSVC is worth the annoyance of saving a couple of
// extra characters on each array access in the implementation code.

// Users of this library can define LAY_REALLOC to use a custom (re)allocator
// instead of stdlib's realloc. It should have the same behavior as realloc --
// first parameter type is a func pointer, and its value is either a null
// pointer or an existing pointer. The second parameter is a size_t of the new
// desired size. The buffer contents should be preserved across reallocations.
//
// And, if you define LAY_REALLOC, you will also need to define LAY_FREE, which
// should have the same behavior as free.

// Like the LAY_REALLOC define, LAY_MEMSET can be used for a custom memset.
// Otherwise, the memset from string.h will be used.

// Useful math utilities
func lay_scalar_max(x, y lay_scalar) lay_scalar {
	return lay_scalar(math.Max(float64(x), float64(y)))
}
func lay_scalar_min(x, y lay_scalar) lay_scalar {
	return lay_scalar(math.Min(float64(x), float64(y)))
}
func lay_float_max(a, b float32) float32 {
	return float32(math.Max(float64(a), float64(b)))
}
func lay_float_min(a, b float32) float32 {
	return float32(math.Min(float64(a), float64(b)))
}

// Performs the layout calculations, starting at the root item (id 0). After
// calling this, you can use lay_get_rect() to query for an item's calculated
// rectangle. If you use procedures such as lay_append() or lay_insert() after
// calling this, your calculated data may become invalid if a reallocation
// occurs.
//
// You should prefer to recreate your items starting from the root instead of
// doing fine-grained updates to the existing context.
//
// However, it's safe to use lay_set_size on an item, and then re-run
// lay_run_context. This might be useful if you are doing a resizing animation
// on items in a layout without any contents changing.
func (ctx *lay_context) lay_run_context() {
	LAY_ASSERT(ctx != nil)

	if ctx.items != nil {
		ctx.lay_run_item(0)
	}
}

// Like lay_run_context(), this procedure will run layout calculations --
// however, it lets you specify which item you want to start from.
// lay_run_context() always starts with item 0, the first item, as the root.
// Running the layout calculations from a specific item is useful if you want
// need to iteratively re-run parts of your layout hierarchy, or if you are only
// interested in updating certain subsets of it. Be careful when using this --
// it's easy to generated bad output if the parent items haven't yet had their
// output rectangles calculated, or if they've been invalidated (e.g. due to
// re-allocation).
func (ctx *lay_context) lay_run_item(item lay_id) {
	LAY_ASSERT(ctx != nil)

	ctx.lay_calc_size(item, 0)
	ctx.lay_arrange(item, 0)
	ctx.lay_calc_size(item, 1)
	ctx.lay_arrange(item, 1)
}

// Alternatively, we could use a flag bit to indicate whether an item's children
// have already been wrapped and may need re-wrapping. If we do that, in the
// future, this would become deprecated and we could make it a no-op.

// Performing a layout on items where wrapping is enabled in the parent
// container can cause flags to be modified during the calculations. If you plan
// to call lay_run_context or lay_run_item multiple times without calling
// lay_reset, and if you have a container that uses wrapping, and if the width
// or height of the container may have changed, you should call
// lay_clear_item_break on all of the children of a container before calling
// lay_run_context or lay_run_item again. If you don't, the layout calculations
// may perform unnecessary wrapping.
//
// This requirement may be changed in the future.
//
// Calling this will also reset any manually-specified breaking. You will need
// to set the manual breaking again, or simply not call this on any items that
// you know you wanted to break manually.
//
// If you clear your context every time you calculate your layout, or if you
// don't use wrapping, you don't need to call this.

func (ctx *lay_context) lay_clear_item_break(item lay_id) {
	LAY_ASSERT(ctx != nil)
	pitem := ctx.lay_get_item(item)
	pitem.flags = pitem.flags & ^LAY_BREAK
}

// Returns the number of items that have been created in a context.
func (ctx *lay_context) lay_items_count() lay_id {
	LAY_ASSERT(ctx != nil)
	return lay_id(len(ctx.items))
}

// Returns the number of items the context can hold without performing a
// reallocation.
func (ctx *lay_context) lay_items_capacity() lay_id {
	LAY_ASSERT(ctx != nil)
	return lay_id(cap(ctx.items))
}

// Create a new item, which can just be thought of as a rectangle. Returns the
// id (handle) used to identify the item.
func (ctx *lay_context) lay_item() lay_id {
	ctx.items = append(ctx.items, &lay_item_t{
		first_child:  LAY_INVALID_ID,
		next_sibling: LAY_INVALID_ID,
	})
	ctx.rects = append(ctx.rects, lay_vec4{})
	return lay_id(len(ctx.items) - 1)
}

func lay_append_by_ptr(pearlier *lay_item_t, later lay_id, plater *lay_item_t) {
	plater.next_sibling = pearlier.next_sibling
	plater.flags |= LAY_ITEM_INSERTED
	pearlier.next_sibling = later
}

func (ctx *lay_context) lay_last_child(parent lay_id) lay_id {
	pparent := ctx.lay_get_item(parent)
	child := pparent.first_child
	if child == LAY_INVALID_ID {
		return LAY_INVALID_ID
	}
	pchild := ctx.lay_get_item(child)
	result := child
	for true {
		next := pchild.next_sibling
		if next == LAY_INVALID_ID {
			break
		}
		result = next
		pchild = ctx.lay_get_item(next)
	}
	return result
}

// lay_append inserts an item as a sibling after another item. This allows
// inserting an item into the middle of an existing list of items within a
// parent. It's also more efficient than repeatedly using lay_insert(ctx,
// parent, new_child) in a loop to create a list of items in a parent, because
// it does not need to traverse the parent's children each time. So if you're
// creating a long list of children inside of a parent, you might prefer to use
// this after using lay_insert to insert the first child.
func (ctx *lay_context) lay_append(earlier lay_id, later lay_id) {
	LAY_ASSERT(later != 0)       // Must not be root item
	LAY_ASSERT(earlier != later) // Must not be same item id
	pearlier := ctx.lay_get_item(earlier)
	plater := ctx.lay_get_item(later)
	lay_append_by_ptr(pearlier, later, plater)
}

// Inserts an item into another item, forming a parent - child relationship. An
// item can contain any number of child items. Items inserted into a parent are
// put at the end of the ordering, after any existing siblings.
func (ctx *lay_context) lay_insert(parent lay_id, child lay_id) {
	LAY_ASSERT(child != 0)      // Must not be root item
	LAY_ASSERT(parent != child) // Must not be same item id
	pparent := ctx.lay_get_item(parent)
	pchild := ctx.lay_get_item(child)
	LAY_ASSERT((pchild.flags & LAY_ITEM_INSERTED) == 0)
	// Parent has no existing children, make inserted item the first child.
	if pparent.first_child == LAY_INVALID_ID {
		pparent.first_child = child
		pchild.flags |= LAY_ITEM_INSERTED
		// Parent has existing items, iterate to find the last child and append the
		// inserted item after it.
	} else {
		next := pparent.first_child
		pnext := ctx.lay_get_item(next)
		for true {
			next = pnext.next_sibling
			if next == LAY_INVALID_ID {
				break
			}
			pnext = ctx.lay_get_item(next)
		}
		lay_append_by_ptr(pnext, child, pchild)
	}
}

// Like lay_insert, but puts the new item as the first child in a parent instead
// of as the last.
func (ctx *lay_context) lay_push(parent lay_id, new_child lay_id) {
	LAY_ASSERT(new_child != 0)      // Must not be root item
	LAY_ASSERT(parent != new_child) // Must not be same item id
	pparent := ctx.lay_get_item(parent)
	old_child := pparent.first_child
	pchild := ctx.lay_get_item(new_child)
	LAY_ASSERT((pchild.flags & LAY_ITEM_INSERTED) == 0)
	pparent.first_child = new_child
	pchild.flags |= LAY_ITEM_INSERTED
	pchild.next_sibling = old_child
}

// Gets the size that was set with lay_set_size or lay_set_size_xy. The _xy
// version writes the output values to the specified addresses instead of
// returning the values in a lay_vec2.
func (ctx *lay_context) lay_get_size(item lay_id) lay_vec2 {
	pitem := ctx.lay_get_item(item)
	return pitem.size
}

func (ctx *lay_context) lay_get_size_xy(item lay_id) lay_vec2 {
	pitem := ctx.lay_get_item(item)
	return pitem.size
}

// Sets the size of an item. The _xy version passes the width and height as
// separate arguments, but functions the same.
func (ctx *lay_context) lay_set_size(item lay_id, size lay_vec2) {
	pitem := ctx.lay_get_item(item)
	pitem.size = size
	flags := pitem.flags
	if size[0] == 0 {
		flags &= ^LAY_ITEM_HFIXED
	} else {
		flags |= LAY_ITEM_HFIXED
	}
	if size[1] == 0 {
		flags &= ^LAY_ITEM_VFIXED
	} else {
		flags |= LAY_ITEM_VFIXED
	}
	pitem.flags = flags
}

func (ctx *lay_context) lay_set_size_xy(item lay_id, width, height lay_scalar) {
	pitem := ctx.lay_get_item(item)
	pitem.size[0] = width
	pitem.size[1] = height
	// Kinda redundant, whatever
	flags := pitem.flags
	if width == 0 {
		flags &= ^LAY_ITEM_HFIXED
	} else {
		flags |= LAY_ITEM_HFIXED
	}
	if height == 0 {
		flags &= ^LAY_ITEM_VFIXED
	} else {
		flags |= LAY_ITEM_VFIXED
	}
	pitem.flags = flags
}

// Set the flags on an item which determines how it behaves as a child inside of
// a parent item. For example, setting LAY_VFILL will make an item try to fill
// up all available vertical space inside of its parent.
func (ctx *lay_context) lay_set_behave(item lay_id, flags int32) {
	LAY_ASSERT((flags & LAY_ITEM_LAYOUT_MASK) == flags)
	pitem := ctx.lay_get_item(item)
	pitem.flags = (pitem.flags & ^LAY_ITEM_LAYOUT_MASK) | flags
}

// Set the flags on an item which determines how it behaves as a parent. For
// example, setting LAY_COLUMN will make an item behave as if it were a column
// -- it will lay out its children vertically.
func (ctx *lay_context) lay_set_contain(item lay_id, flags int32) {
	LAY_ASSERT((flags & LAY_ITEM_BOX_MASK) == flags)
	pitem := ctx.lay_get_item(item)
	pitem.flags = (pitem.flags & ^LAY_ITEM_BOX_MASK) | flags
}

// Set the margins on an item. The components of the vector are:
// 0: left, 1: top, 2: right, 3: bottom.
func (ctx *lay_context) lay_set_margins(item lay_id, ltrb lay_vec4) {
	pitem := ctx.lay_get_item(item)
	pitem.margins = ltrb
}

// Same as lay_set_margins, but the components are passed as separate arguments
// (left, top, right, bottom).
func (ctx *lay_context) lay_set_margins_ltrb(item lay_id, l, t, r, b lay_scalar) {
	pitem := ctx.lay_get_item(item)
	// Alternative, uses stack and addressed writes
	//pitem.margins = lay_vec4_xyzw(l, t, r, b)
	// Alternative, uses rax and left-shift
	//pitem.margins = (lay_vec4){l, t, r, b}
	// Fewest instructions, but uses more addressed writes?
	pitem.margins[0] = l
	pitem.margins[1] = t
	pitem.margins[2] = r
	pitem.margins[3] = b
}

// Get the margins that were set by lay_set_margins. The _ltrb version writes
// the output values to the specified addresses instead of returning the values
// in a lay_vec4.
// l: left, t: top, r: right, b: bottom
func (ctx *lay_context) lay_get_margins(item lay_id) lay_vec4 {
	return ctx.lay_get_item(item).margins
}

func (ctx *lay_context) lay_get_margins_ltrb(item lay_id) lay_vec4 {
	pitem := ctx.lay_get_item(item)
	return pitem.margins
}

// TODO restrict item ptrs correctly
func (ctx *lay_context) lay_calc_overlayed_size(item lay_id, dim int) lay_scalar {
	wdim := dim + 2
	pitem := ctx.lay_get_item(item)
	var need_size lay_scalar
	child := pitem.first_child
	for child != LAY_INVALID_ID {
		pchild := ctx.lay_get_item(child)
		rect := ctx.rects[child]
		// width = start margin + calculated width + end margin
		child_size := rect[dim] + rect[2+dim] + pchild.margins[wdim]
		need_size = lay_scalar_max(need_size, child_size)
		child = pchild.next_sibling
	}
	return need_size
}

func (ctx *lay_context) lay_calc_stacked_size(item lay_id, dim int) lay_scalar {
	wdim := dim + 2
	pitem := ctx.lay_get_item(item)
	var need_size lay_scalar
	child := pitem.first_child
	for child != LAY_INVALID_ID {
		pchild := ctx.lay_get_item(child)
		rect := ctx.rects[child]
		need_size += rect[dim] + rect[2+dim] + pchild.margins[wdim]
		child = pchild.next_sibling
	}
	return need_size
}

func (ctx *lay_context) lay_calc_wrapped_overlayed_size(item lay_id, dim int) lay_scalar {
	wdim := dim + 2
	pitem := ctx.lay_get_item(item)
	var need_size lay_scalar
	var need_size2 lay_scalar
	child := pitem.first_child
	for child != LAY_INVALID_ID {
		pchild := ctx.lay_get_item(child)
		rect := ctx.rects[child]
		if (pchild.flags & LAY_BREAK) != 0 {
			need_size2 += need_size
			need_size = 0
		}
		child_size := rect[dim] + rect[2+dim] + pchild.margins[wdim]
		need_size = lay_scalar_max(need_size, child_size)
		child = pchild.next_sibling
	}
	return need_size2 + need_size
}

// Equivalent to uiComputeWrappedStackedSize
func (ctx *lay_context) lay_calc_wrapped_stacked_size(item lay_id, dim int) lay_scalar {
	wdim := dim + 2
	pitem := ctx.lay_get_item(item)
	var need_size lay_scalar
	var need_size2 lay_scalar
	child := pitem.first_child
	for child != LAY_INVALID_ID {
		pchild := ctx.lay_get_item(child)
		rect := ctx.rects[child]
		if (pchild.flags & LAY_BREAK) != 0 {
			need_size2 = lay_scalar_max(need_size2, need_size)
			need_size = 0
		}
		need_size += rect[dim] + rect[2+dim] + pchild.margins[wdim]
		child = pchild.next_sibling
	}
	return lay_scalar_max(need_size2, need_size)
}

func (ctx *lay_context) lay_calc_size(item lay_id, dim int) {
	pitem := ctx.lay_get_item(item)

	child := pitem.first_child
	for child != LAY_INVALID_ID {
		// NOTE: this is recursive and will run out of stack space if items are
		// nested too deeply.
		ctx.lay_calc_size(child, dim)
		pchild := ctx.lay_get_item(child)
		child = pchild.next_sibling
	}

	// Set the mutable rect output data to the starting input data
	ctx.rects[item][dim] = pitem.margins[dim]

	// If we have an explicit input size, just set our output size (which other
	// calc_size and arrange procedures will use) to it.
	if pitem.size[dim] != 0 {
		ctx.rects[item][2+dim] = pitem.size[dim]
		return
	}

	// Calculate our size based on children items. Note that we've already
	// called lay_calc_size on our children at this point.
	var cal_size lay_scalar
	switch pitem.flags & LAY_ITEM_BOX_MODEL_MASK {
	case LAY_COLUMN | LAY_WRAP:
		// flex model
		if dim != 0 { // direction
			cal_size = ctx.lay_calc_stacked_size(item, 1)
		} else {
			cal_size = ctx.lay_calc_overlayed_size(item, 0)
		}
		break
	case LAY_ROW | LAY_WRAP:
		// flex model
		if dim == 0 { // direction
			cal_size = ctx.lay_calc_wrapped_stacked_size(item, 0)
		} else {
			cal_size = ctx.lay_calc_wrapped_overlayed_size(item, 1)
		}
		break
	case LAY_COLUMN:
	case LAY_ROW:
		// flex model
		if (pitem.flags & 1) == int32(dim) { // direction
			cal_size = ctx.lay_calc_stacked_size(item, dim)
		} else {
			cal_size = ctx.lay_calc_overlayed_size(item, dim)
		}
		break
	default:
		// layout model
		cal_size = ctx.lay_calc_overlayed_size(item, dim)
		break
	}

	// Set our output data size. Will be used by parent calc_size procedures.,
	// and by arrange procedures.
	ctx.rects[item][2+dim] = cal_size
}

func (ctx *lay_context) lay_arrange_stacked(item lay_id, dim int, wrap bool) {
	wdim := dim + 2
	pitem := ctx.lay_get_item(item)

	item_flags := pitem.flags
	rect := ctx.rects[item]
	space := rect[2+dim]

	max_x2 := float32(rect[dim] + space)

	start_child := pitem.first_child
	for start_child != LAY_INVALID_ID {
		var used lay_scalar
		var count int          // count of fillers
		var squeezed_count int // count of squeezable elements
		var total int
		var hardbreak bool
		// first pass: count items that need to be expanded,
		// and the space that is used
		child := start_child
		end_child := LAY_INVALID_ID
		for child != LAY_INVALID_ID {
			pchild := ctx.lay_get_item(child)
			child_flags := pchild.flags
			flags := (child_flags & LAY_ITEM_LAYOUT_MASK) >> uint(dim)
			fflags := (child_flags & LAY_ITEM_FIXED_MASK) >> uint(dim)
			child_margins := pchild.margins
			child_rect := ctx.rects[child]
			extend := used
			if (flags & LAY_HFILL) == LAY_HFILL {
				count++
				extend += child_rect[dim] + child_margins[wdim]
			} else {
				if (fflags & LAY_ITEM_HFIXED) != LAY_ITEM_HFIXED {
					squeezed_count++
				}
				extend += child_rect[dim] + child_rect[2+dim] + child_margins[wdim]
			}
			// wrap on end of line or manual flag
			if wrap && (total != 0 && ((extend > space) ||
				(child_flags&LAY_BREAK) != 0)) {
				end_child = child
				hardbreak = (child_flags & LAY_BREAK) == LAY_BREAK
				// add marker for subsequent queries
				pchild.flags = child_flags | LAY_BREAK
				break
			} else {
				used = extend
				child = pchild.next_sibling
			}
			total++
		}

		extra_space := space - used
		var filler float32
		var spacer float32
		var extra_margin float32
		var eater float32

		if extra_space > 0 {
			if count > 0 {
				filler = float32(extra_space) / float32(count)
			} else if total > 0 {
				switch item_flags & LAY_JUSTIFY {
				case LAY_JUSTIFY:
					// justify when not wrapping or not in last line,
					// or not manually breaking
					if !wrap || ((end_child != LAY_INVALID_ID) && !hardbreak) {
						spacer = float32(extra_space) / float32(total-1)
					}
					break
				case LAY_START:
					break
				case LAY_END:
					extra_margin = float32(extra_space)
					break
				default:
					extra_margin = float32(extra_space) / 2.0
					break
				}
			}
			//#ifdef LAY_FLOAT
			// In floating point, it's possible to end up with some small negative
			// value for extra_space, while also have a 0.0 squeezed_count. This
			// would cause divide by zero. Instead, we'll check to see if
			// squeezed_count is > 0. I believe this produces the same results as
			// the original oui int-only code. However, I don't have any tests for
			// it, so I'll leave it if-def'd for now.
			//        else if (!wrap && (squeezed_count > 0))
			//#else
			// This is the original oui code
		} else if !wrap && (extra_space < 0) {
			//#endif
			eater = float32(extra_space) / float32(squeezed_count)
		}
		// distribute width among items
		x := float32(rect[dim])
		x1 := float32(0)
		// second pass: distribute and rescale
		child = start_child
		for child != end_child {
			var ix0, ix1 lay_scalar
			pchild := ctx.lay_get_item(child)
			child_flags := pchild.flags
			flags := (child_flags & LAY_ITEM_LAYOUT_MASK) >> uint(dim)
			fflags := (child_flags & LAY_ITEM_FIXED_MASK) >> uint(dim)
			child_margins := pchild.margins
			child_rect := ctx.rects[child]

			x += float32(child_rect[dim]) + extra_margin
			if (flags & LAY_HFILL) == LAY_HFILL { // grow
				x1 = x + filler
			} else if (fflags & LAY_ITEM_HFIXED) == LAY_ITEM_HFIXED {
				x1 = x + float32(child_rect[2+dim])
			} else { // squeeze
				x1 = x + lay_float_max(0.0, float32(child_rect[2+dim])+eater)
			}
			ix0 = lay_scalar(x)
			if wrap {
				ix1 = lay_scalar(lay_float_min(max_x2-float32(child_margins[wdim]), x1))
			} else {
				ix1 = lay_scalar(x1)
			}
			child_rect[dim] = ix0         // pos
			child_rect[dim+2] = ix1 - ix0 // size
			ctx.rects[child] = child_rect
			x = x1 + float32(child_margins[wdim])
			child = pchild.next_sibling
			extra_margin = spacer
		}

		start_child = end_child
	}
}

func (ctx *lay_context) lay_arrange_overlay(item lay_id, dim int) {
	wdim := dim + 2
	pitem := ctx.lay_get_item(item)
	rect := ctx.rects[item]
	offset := rect[dim]
	space := rect[2+dim]

	child := pitem.first_child
	for child != LAY_INVALID_ID {
		pchild := ctx.lay_get_item(child)
		b_flags := (pchild.flags & LAY_ITEM_LAYOUT_MASK) >> uint(dim)
		child_margins := pchild.margins
		child_rect := ctx.rects[child]

		switch b_flags & LAY_HFILL {
		case LAY_HCENTER:
			child_rect[dim] += (space-child_rect[2+dim])/2 - child_margins[wdim]
			break
		case LAY_RIGHT:
			child_rect[dim] += space - child_rect[2+dim] - child_margins[wdim]
			break
		case LAY_HFILL:
			child_rect[2+dim] = lay_scalar_max(0, space-child_rect[dim]-child_margins[wdim])
			break
		default:
			break
		}

		child_rect[dim] += offset
		ctx.rects[child] = child_rect
		child = pchild.next_sibling
	}
}

func (ctx *lay_context) lay_arrange_overlay_squeezed_range(dim int,
	start_item, end_item lay_id,
	offset, space lay_scalar) {
	wdim := dim + 2
	item := start_item
	for item != end_item {
		pitem := ctx.lay_get_item(item)
		b_flags := (pitem.flags & LAY_ITEM_LAYOUT_MASK) >> uint(dim)
		margins := pitem.margins
		rect := ctx.rects[item]
		min_size := lay_scalar_max(0, space-rect[dim]-margins[wdim])
		switch b_flags & LAY_HFILL {
		case LAY_HCENTER:
			rect[2+dim] = lay_scalar_min(rect[2+dim], min_size)
			rect[dim] += (space-rect[2+dim])/2 - margins[wdim]
			break
		case LAY_RIGHT:
			rect[2+dim] = lay_scalar_min(rect[2+dim], min_size)
			rect[dim] = space - rect[2+dim] - margins[wdim]
			break
		case LAY_HFILL:
			rect[2+dim] = min_size
			break
		default:
			rect[2+dim] = lay_scalar_min(rect[2+dim], min_size)
			break
		}
		rect[dim] += offset
		ctx.rects[item] = rect
		item = pitem.next_sibling
	}
}

func (ctx *lay_context) lay_arrange_wrapped_overlay_squeezed(item lay_id, dim int) lay_scalar {
	wdim := dim + 2
	pitem := ctx.lay_get_item(item)
	offset := ctx.rects[item][dim]
	need_size := lay_scalar(0)
	child := pitem.first_child
	start_child := child
	for child != LAY_INVALID_ID {
		pchild := ctx.lay_get_item(child)
		if (pchild.flags & LAY_BREAK) != 0 {
			ctx.lay_arrange_overlay_squeezed_range(dim, start_child, child, offset, need_size)
			offset += need_size
			start_child = child
			need_size = 0
		}
		rect := ctx.rects[child]
		child_size := rect[dim] + rect[2+dim] + pchild.margins[wdim]
		need_size = lay_scalar_max(need_size, child_size)
		child = pchild.next_sibling
	}
	ctx.lay_arrange_overlay_squeezed_range(dim, start_child, LAY_INVALID_ID, offset, need_size)
	offset += need_size
	return offset
}

func (ctx *lay_context) lay_arrange(item lay_id, dim int) {
	pitem := ctx.lay_get_item(item)

	flags := pitem.flags
	switch flags & LAY_ITEM_BOX_MODEL_MASK {
	case LAY_COLUMN | LAY_WRAP:
		if dim != 0 {
			ctx.lay_arrange_stacked(item, 1, true)
			offset := ctx.lay_arrange_wrapped_overlay_squeezed(item, 0)
			ctx.rects[item][2+0] = offset - ctx.rects[item][0]
		}
		break
	case LAY_ROW | LAY_WRAP:
		if dim == 0 {
			ctx.lay_arrange_stacked(item, 0, true)
		} else {
			// discard return value
			ctx.lay_arrange_wrapped_overlay_squeezed(item, 1)
		}
		break
	case LAY_COLUMN:
	case LAY_ROW:
		if (flags & 1) == int32(dim) {
			ctx.lay_arrange_stacked(item, dim, false)
		} else {
			rect := ctx.rects[item]
			ctx.lay_arrange_overlay_squeezed_range(
				dim, pitem.first_child, LAY_INVALID_ID,
				rect[dim], rect[2+dim])
		}
		break
	default:
		ctx.lay_arrange_overlay(item, dim)
		break
	}
	child := pitem.first_child
	for child != LAY_INVALID_ID {
		// NOTE: this is recursive and will run out of stack space if items are
		// nested too deeply.
		ctx.lay_arrange(child, dim)
		pchild := ctx.lay_get_item(child)
		child = pchild.next_sibling
	}
}
