package goquery

import (
	"strings"
	"testing"

	"code.google.com/p/cascadia"
)

func TestFind(t *testing.T) {
	sel := Doc().Find("div.row-fluid")
	AssertLength(t, sel.Nodes, 9)

	selc := Doc().FindCompiled(cascadia.MustCompile("div.row-fluid"))
	AssertSelEqual(t, sel, selc)
}

func TestFindRollback(t *testing.T) {
	sel := Doc().Find("div.row-fluid")
	sel2 := sel.Find("a").End()
	AssertEqual(t, sel, sel2)
}

func TestFindNotSelf(t *testing.T) {
	sel := Doc().Find("h1").Find("h1")
	AssertLength(t, sel.Nodes, 0)
}

func TestFindInvalidSelector(t *testing.T) {
	defer AssertPanic(t)
	Doc().Find(":+ ^")
}

func TestChainedFind(t *testing.T) {
	sel := Doc().Find("div.hero-unit").Find(".row-fluid")
	AssertLength(t, sel.Nodes, 4)
}

func TestChildren(t *testing.T) {
	sel := Doc().Find(".pvk-content").Children()
	AssertLength(t, sel.Nodes, 5)
}

func TestChildrenRollback(t *testing.T) {
	sel := Doc().Find(".pvk-content")
	sel2 := sel.Children().End()
	AssertEqual(t, sel, sel2)
}

func TestContents(t *testing.T) {
	sel := Doc().Find(".pvk-content").Contents()
	AssertLength(t, sel.Nodes, 13)
}

func TestContentsRollback(t *testing.T) {
	sel := Doc().Find(".pvk-content")
	sel2 := sel.Contents().End()
	AssertEqual(t, sel, sel2)
}

func TestChildrenFiltered(t *testing.T) {
	sel := Doc().Find(".pvk-content").ChildrenFiltered(".hero-unit")
	AssertLength(t, sel.Nodes, 1)

	selc := Doc().Find(".pvk-content").ChildrenFilteredCompiled(cascadia.MustCompile(".hero-unit"))
	AssertSelEqual(t, sel, selc)
}

func TestChildrenFilteredRollback(t *testing.T) {
	sel := Doc().Find(".pvk-content")
	sel2 := sel.ChildrenFiltered(".hero-unit").End()
	AssertEqual(t, sel, sel2)
}

func TestContentsFiltered(t *testing.T) {
	sel := Doc().Find(".pvk-content").ContentsFiltered(".hero-unit")
	AssertLength(t, sel.Nodes, 1)

	selc := Doc().Find(".pvk-content").ContentsFilteredCompiled(cascadia.MustCompile(".hero-unit"))
	AssertSelEqual(t, sel, selc)
}

func TestContentsFilteredRollback(t *testing.T) {
	sel := Doc().Find(".pvk-content")
	sel2 := sel.ContentsFiltered(".hero-unit").End()
	AssertEqual(t, sel, sel2)
}

func TestChildrenFilteredNone(t *testing.T) {
	sel := Doc().Find(".pvk-content").ChildrenFiltered("a.btn")
	AssertLength(t, sel.Nodes, 0)
}

func TestParent(t *testing.T) {
	sel := Doc().Find(".container-fluid").Parent()
	AssertLength(t, sel.Nodes, 3)
}

func TestParentRollback(t *testing.T) {
	sel := Doc().Find(".container-fluid")
	sel2 := sel.Parent().End()
	AssertEqual(t, sel, sel2)
}

func TestParentBody(t *testing.T) {
	sel := Doc().Find("body").Parent()
	AssertLength(t, sel.Nodes, 1)
}

func TestParentFiltered(t *testing.T) {
	sel := Doc().Find(".container-fluid").ParentFiltered(".hero-unit")
	AssertLength(t, sel.Nodes, 1)
	AssertClass(t, sel, "hero-unit")

	selc := Doc().Find(".container-fluid").ParentFilteredCompiled(cascadia.MustCompile(".hero-unit"))
	AssertSelEqual(t, sel, selc)
}

func TestParentFilteredRollback(t *testing.T) {
	sel := Doc().Find(".container-fluid")
	sel2 := sel.ParentFiltered(".hero-unit").End()
	AssertEqual(t, sel, sel2)
}

func TestParents(t *testing.T) {
	sel := Doc().Find(".container-fluid").Parents()
	AssertLength(t, sel.Nodes, 8)
}

func TestParentsOrder(t *testing.T) {
	sel := Doc().Find("#cf2").Parents()
	AssertLength(t, sel.Nodes, 6)
	AssertSelectionIs(t, sel, ".hero-unit", ".pvk-content", "div.row-fluid", "#cf1", "body", "html")
}

func TestParentsRollback(t *testing.T) {
	sel := Doc().Find(".container-fluid")
	sel2 := sel.Parents().End()
	AssertEqual(t, sel, sel2)
}

func TestParentsFiltered(t *testing.T) {
	sel := Doc().Find(".container-fluid").ParentsFiltered("body")
	AssertLength(t, sel.Nodes, 1)

	selc := Doc().Find(".container-fluid").ParentsFilteredCompiled(cascadia.MustCompile("body"))
	AssertSelEqual(t, sel, selc)
}

func TestParentsFilteredRollback(t *testing.T) {
	sel := Doc().Find(".container-fluid")
	sel2 := sel.ParentsFiltered("body").End()
	AssertEqual(t, sel, sel2)
}

func TestParentsUntil(t *testing.T) {
	sel := Doc().Find(".container-fluid").ParentsUntil("body")
	AssertLength(t, sel.Nodes, 6)

	selc := Doc().Find(".container-fluid").ParentsUntilCompiled(cascadia.MustCompile("body"))
	AssertSelEqual(t, sel, selc)
}

func TestParentsUntilRollback(t *testing.T) {
	sel := Doc().Find(".container-fluid")
	sel2 := sel.ParentsUntil("body").End()
	AssertEqual(t, sel, sel2)
}

func TestParentsUntilSelection(t *testing.T) {
	sel := Doc().Find(".container-fluid")
	sel2 := Doc().Find(".pvk-content")
	sel = sel.ParentsUntilSelection(sel2)
	AssertLength(t, sel.Nodes, 3)
}

func TestParentsUntilSelectionRollback(t *testing.T) {
	sel := Doc().Find(".container-fluid")
	sel2 := Doc().Find(".pvk-content")
	sel2 = sel.ParentsUntilSelection(sel2).End()
	AssertEqual(t, sel, sel2)
}

func TestParentsUntilNodes(t *testing.T) {
	sel := Doc().Find(".container-fluid")
	sel2 := Doc().Find(".pvk-content, .hero-unit")
	sel = sel.ParentsUntilNodes(sel2.Nodes...)
	AssertLength(t, sel.Nodes, 2)
}

func TestParentsUntilNodesRollback(t *testing.T) {
	sel := Doc().Find(".container-fluid")
	sel2 := Doc().Find(".pvk-content, .hero-unit")
	sel2 = sel.ParentsUntilNodes(sel2.Nodes...).End()
	AssertEqual(t, sel, sel2)
}

func TestParentsFilteredUntil(t *testing.T) {
	sel := Doc().Find(".container-fluid").ParentsFilteredUntil(".pvk-content", "body")
	AssertLength(t, sel.Nodes, 2)

	selc := Doc().Find(".container-fluid").ParentsFilteredUntilCompiled(
		cascadia.MustCompile(".pvk-content"),
		cascadia.MustCompile("body"))
	AssertSelEqual(t, sel, selc)
}

func TestParentsFilteredUntilRollback(t *testing.T) {
	sel := Doc().Find(".container-fluid")
	sel2 := sel.ParentsFilteredUntil(".pvk-content", "body").End()
	AssertEqual(t, sel, sel2)
}

func TestParentsFilteredUntilSelection(t *testing.T) {
	sel := Doc().Find(".container-fluid")
	sel2 := Doc().Find(".row-fluid")
	sel3 := sel.ParentsFilteredUntilSelection("div", sel2)
	AssertLength(t, sel3.Nodes, 3)

	selc := sel.ParentsFilteredUntilSelectionCompiled(cascadia.MustCompile("div"), sel2)
	AssertSelEqual(t, sel3, selc)
}

func TestParentsFilteredUntilSelectionRollback(t *testing.T) {
	sel := Doc().Find(".container-fluid")
	sel2 := Doc().Find(".row-fluid")
	sel2 = sel.ParentsFilteredUntilSelection("div", sel2).End()
	AssertEqual(t, sel, sel2)
}

func TestParentsFilteredUntilNodes(t *testing.T) {
	sel := Doc().Find(".container-fluid")
	sel2 := Doc().Find(".row-fluid")
	sel = sel.ParentsFilteredUntilNodes("body", sel2.Nodes...)
	AssertLength(t, sel.Nodes, 1)
}

func TestParentsFilteredUntilNodesRollback(t *testing.T) {
	sel := Doc().Find(".container-fluid")
	sel2 := Doc().Find(".row-fluid")
	sel2 = sel.ParentsFilteredUntilNodes("body", sel2.Nodes...).End()
	AssertEqual(t, sel, sel2)
}

func TestSiblings(t *testing.T) {
	sel := Doc().Find("h1").Siblings()
	AssertLength(t, sel.Nodes, 1)
}

func TestSiblingsRollback(t *testing.T) {
	sel := Doc().Find("h1")
	sel2 := sel.Siblings().End()
	AssertEqual(t, sel, sel2)
}

func TestSiblings2(t *testing.T) {
	sel := Doc().Find(".pvk-gutter").Siblings()
	AssertLength(t, sel.Nodes, 9)
}

func TestSiblings3(t *testing.T) {
	sel := Doc().Find("body>.container-fluid").Siblings()
	AssertLength(t, sel.Nodes, 0)
}

func TestSiblingsFiltered(t *testing.T) {
	sel := Doc().Find(".pvk-gutter").SiblingsFiltered(".pvk-content")
	AssertLength(t, sel.Nodes, 3)

	selc := Doc().Find(".pvk-gutter").SiblingsFilteredCompiled(cascadia.MustCompile(".pvk-content"))
	AssertSelEqual(t, sel, selc)
}

func TestSiblingsFilteredRollback(t *testing.T) {
	sel := Doc().Find(".pvk-gutter")
	sel2 := sel.SiblingsFiltered(".pvk-content").End()
	AssertEqual(t, sel, sel2)
}

func TestNext(t *testing.T) {
	sel := Doc().Find("h1").Next()
	AssertLength(t, sel.Nodes, 1)
}

func TestNextRollback(t *testing.T) {
	sel := Doc().Find("h1")
	sel2 := sel.Next().End()
	AssertEqual(t, sel, sel2)
}

func TestNext2(t *testing.T) {
	sel := Doc().Find(".close").Next()
	AssertLength(t, sel.Nodes, 1)
}

func TestNextNone(t *testing.T) {
	sel := Doc().Find("small").Next()
	AssertLength(t, sel.Nodes, 0)
}

func TestNextFiltered(t *testing.T) {
	sel := Doc().Find(".container-fluid").NextFiltered("div")
	AssertLength(t, sel.Nodes, 2)

	selc := Doc().Find(".container-fluid").NextFilteredCompiled(cascadia.MustCompile("div"))
	AssertSelEqual(t, sel, selc)
}

func TestNextFilteredRollback(t *testing.T) {
	sel := Doc().Find(".container-fluid")
	sel2 := sel.NextFiltered("div").End()
	AssertEqual(t, sel, sel2)
}

func TestNextFiltered2(t *testing.T) {
	sel := Doc().Find(".container-fluid").NextFiltered("[ng-view]")
	AssertLength(t, sel.Nodes, 1)

	selc := Doc().Find(".container-fluid").NextFilteredCompiled(cascadia.MustCompile("[ng-view]"))
	AssertSelEqual(t, sel, selc)
}

func TestPrev(t *testing.T) {
	sel := Doc().Find(".red").Prev()
	AssertLength(t, sel.Nodes, 1)
	AssertClass(t, sel, "green")
}

func TestPrevRollback(t *testing.T) {
	sel := Doc().Find(".red")
	sel2 := sel.Prev().End()
	AssertEqual(t, sel, sel2)
}

func TestPrev2(t *testing.T) {
	sel := Doc().Find(".row-fluid").Prev()
	AssertLength(t, sel.Nodes, 5)
}

func TestPrevNone(t *testing.T) {
	sel := Doc().Find("h2").Prev()
	AssertLength(t, sel.Nodes, 0)
}

func TestPrevFiltered(t *testing.T) {
	sel := Doc().Find(".row-fluid").PrevFiltered(".row-fluid")
	AssertLength(t, sel.Nodes, 5)

	selc := Doc().Find(".row-fluid").PrevFilteredCompiled(cascadia.MustCompile(".row-fluid"))
	AssertSelEqual(t, sel, selc)
}

func TestPrevFilteredRollback(t *testing.T) {
	sel := Doc().Find(".row-fluid")
	sel2 := sel.PrevFiltered(".row-fluid").End()
	AssertEqual(t, sel, sel2)
}

func TestNextAll(t *testing.T) {
	sel := Doc().Find("#cf2 div:nth-child(1)").NextAll()
	AssertLength(t, sel.Nodes, 3)
}

func TestNextAllRollback(t *testing.T) {
	sel := Doc().Find("#cf2 div:nth-child(1)")
	sel2 := sel.NextAll().End()
	AssertEqual(t, sel, sel2)
}

func TestNextAll2(t *testing.T) {
	sel := Doc().Find("div[ng-cloak]").NextAll()
	AssertLength(t, sel.Nodes, 1)
}

func TestNextAllNone(t *testing.T) {
	sel := Doc().Find(".footer").NextAll()
	AssertLength(t, sel.Nodes, 0)
}

func TestNextAllFiltered(t *testing.T) {
	sel := Doc().Find("#cf2 .row-fluid").NextAllFiltered("[ng-cloak]")
	AssertLength(t, sel.Nodes, 2)

	selc := Doc().Find("#cf2 .row-fluid").NextAllFilteredCompiled(cascadia.MustCompile("[ng-cloak]"))
	AssertSelEqual(t, sel, selc)
}

func TestNextAllFilteredRollback(t *testing.T) {
	sel := Doc().Find("#cf2 .row-fluid")
	sel2 := sel.NextAllFiltered("[ng-cloak]").End()
	AssertEqual(t, sel, sel2)
}

func TestNextAllFiltered2(t *testing.T) {
	sel := Doc().Find(".close").NextAllFiltered("h4")
	AssertLength(t, sel.Nodes, 1)

	selc := Doc().Find(".close").NextAllFilteredCompiled(cascadia.MustCompile("h4"))
	AssertSelEqual(t, sel, selc)
}

func TestPrevAll(t *testing.T) {
	sel := Doc().Find("[ng-view]").PrevAll()
	AssertLength(t, sel.Nodes, 2)
}

func TestPrevAllOrder(t *testing.T) {
	sel := Doc().Find("[ng-view]").PrevAll()
	AssertLength(t, sel.Nodes, 2)
	AssertSelectionIs(t, sel, "#cf4", "#cf3")
}

func TestPrevAllRollback(t *testing.T) {
	sel := Doc().Find("[ng-view]")
	sel2 := sel.PrevAll().End()
	AssertEqual(t, sel, sel2)
}

func TestPrevAll2(t *testing.T) {
	sel := Doc().Find(".pvk-gutter").PrevAll()
	AssertLength(t, sel.Nodes, 6)
}

func TestPrevAllFiltered(t *testing.T) {
	sel := Doc().Find(".pvk-gutter").PrevAllFiltered(".pvk-content")
	AssertLength(t, sel.Nodes, 3)

	selc := Doc().Find(".pvk-gutter").PrevAllFilteredCompiled(cascadia.MustCompile(".pvk-content"))
	AssertSelEqual(t, sel, selc)
}

func TestPrevAllFilteredRollback(t *testing.T) {
	sel := Doc().Find(".pvk-gutter")
	sel2 := sel.PrevAllFiltered(".pvk-content").End()
	AssertEqual(t, sel, sel2)
}

func TestNextUntil(t *testing.T) {
	sel := Doc().Find(".alert a").NextUntil("p")
	AssertLength(t, sel.Nodes, 1)
	AssertSelectionIs(t, sel, "h4")

	selc := Doc().Find(".alert a").NextUntilCompiled(cascadia.MustCompile("p"))
	AssertSelEqual(t, sel, selc)
}

func TestNextUntil2(t *testing.T) {
	sel := Doc().Find("#cf2-1").NextUntil("[ng-cloak]")
	AssertLength(t, sel.Nodes, 1)
	AssertSelectionIs(t, sel, "#cf2-2")

	selc := Doc().Find("#cf2-1").NextUntilCompiled(cascadia.MustCompile("[ng-cloak]"))
	AssertSelEqual(t, sel, selc)
}

func TestNextUntilOrder(t *testing.T) {
	sel := Doc().Find("#cf2-1").NextUntil("#cf2-4")
	AssertLength(t, sel.Nodes, 2)
	AssertSelectionIs(t, sel, "#cf2-2", "#cf2-3")

	selc := Doc().Find("#cf2-1").NextUntilCompiled(cascadia.MustCompile("#cf2-4"))
	AssertSelEqual(t, sel, selc)
}

func TestNextUntilRollback(t *testing.T) {
	sel := Doc().Find("#cf2-1")
	sel2 := sel.PrevUntil("#cf2-4").End()
	AssertEqual(t, sel, sel2)
}

func TestNextUntilSelection(t *testing.T) {
	sel := Doc2().Find("#n2")
	sel2 := Doc2().Find("#n4")
	sel2 = sel.NextUntilSelection(sel2)
	AssertLength(t, sel2.Nodes, 1)
	AssertSelectionIs(t, sel2, "#n3")
}

func TestNextUntilSelectionRollback(t *testing.T) {
	sel := Doc2().Find("#n2")
	sel2 := Doc2().Find("#n4")
	sel2 = sel.NextUntilSelection(sel2).End()
	AssertEqual(t, sel, sel2)
}

func TestNextUntilNodes(t *testing.T) {
	sel := Doc2().Find("#n2")
	sel2 := Doc2().Find("#n5")
	sel2 = sel.NextUntilNodes(sel2.Nodes...)
	AssertLength(t, sel2.Nodes, 2)
	AssertSelectionIs(t, sel2, "#n3", "#n4")
}

func TestNextUntilNodesRollback(t *testing.T) {
	sel := Doc2().Find("#n2")
	sel2 := Doc2().Find("#n5")
	sel2 = sel.NextUntilNodes(sel2.Nodes...).End()
	AssertEqual(t, sel, sel2)
}

func TestPrevUntil(t *testing.T) {
	sel := Doc().Find(".alert p").PrevUntil("a")
	AssertLength(t, sel.Nodes, 1)
	AssertSelectionIs(t, sel, "h4")

	selc := Doc().Find(".alert p").PrevUntilCompiled(cascadia.MustCompile("a"))
	AssertSelEqual(t, sel, selc)
}

func TestPrevUntil2(t *testing.T) {
	sel := Doc().Find("[ng-cloak]").PrevUntil(":not([ng-cloak])")
	AssertLength(t, sel.Nodes, 1)
	AssertSelectionIs(t, sel, "[ng-cloak]")

	selc := Doc().Find("[ng-cloak]").PrevUntilCompiled(cascadia.MustCompile(":not([ng-cloak])"))
	AssertSelEqual(t, sel, selc)
}

func TestPrevUntilOrder(t *testing.T) {
	sel := Doc().Find("#cf2-4").PrevUntil("#cf2-1")
	AssertLength(t, sel.Nodes, 2)
	AssertSelectionIs(t, sel, "#cf2-3", "#cf2-2")

	selc := Doc().Find("#cf2-4").PrevUntilCompiled(cascadia.MustCompile("#cf2-1"))
	AssertSelEqual(t, sel, selc)
}

func TestPrevUntilRollback(t *testing.T) {
	sel := Doc().Find("#cf2-4")
	sel2 := sel.PrevUntil("#cf2-1").End()
	AssertEqual(t, sel, sel2)
}

func TestPrevUntilSelection(t *testing.T) {
	sel := Doc2().Find("#n4")
	sel2 := Doc2().Find("#n2")
	sel2 = sel.PrevUntilSelection(sel2)
	AssertLength(t, sel2.Nodes, 1)
	AssertSelectionIs(t, sel2, "#n3")
}

func TestPrevUntilSelectionRollback(t *testing.T) {
	sel := Doc2().Find("#n4")
	sel2 := Doc2().Find("#n2")
	sel2 = sel.PrevUntilSelection(sel2).End()
	AssertEqual(t, sel, sel2)
}

func TestPrevUntilNodes(t *testing.T) {
	sel := Doc2().Find("#n5")
	sel2 := Doc2().Find("#n2")
	sel2 = sel.PrevUntilNodes(sel2.Nodes...)
	AssertLength(t, sel2.Nodes, 2)
	AssertSelectionIs(t, sel2, "#n4", "#n3")
}

func TestPrevUntilNodesRollback(t *testing.T) {
	sel := Doc2().Find("#n5")
	sel2 := Doc2().Find("#n2")
	sel2 = sel.PrevUntilNodes(sel2.Nodes...).End()
	AssertEqual(t, sel, sel2)
}

func TestNextFilteredUntil(t *testing.T) {
	sel := Doc2().Find(".two").NextFilteredUntil(".even", ".six")
	AssertLength(t, sel.Nodes, 4)
	AssertSelectionIs(t, sel, "#n3", "#n5", "#nf3", "#nf5")

	selc := Doc2().Find(".two").NextFilteredUntilCompiled(
		cascadia.MustCompile(".even"),
		cascadia.MustCompile(".six"))
	AssertSelEqual(t, sel, selc)
}

func TestNextFilteredUntilRollback(t *testing.T) {
	sel := Doc2().Find(".two")
	sel2 := sel.NextFilteredUntil(".even", ".six").End()
	AssertEqual(t, sel, sel2)
}

func TestNextFilteredUntilSelection(t *testing.T) {
	sel := Doc2().Find(".even")
	sel2 := Doc2().Find(".five")
	sel3 := sel.NextFilteredUntilSelection(".even", sel2)
	AssertLength(t, sel3.Nodes, 2)
	AssertSelectionIs(t, sel3, "#n3", "#nf3")

	selc := sel.NextFilteredUntilSelectionCompiled(cascadia.MustCompile(".even"), sel2)
	AssertSelEqual(t, sel3, selc)
}

func TestNextFilteredUntilSelectionRollback(t *testing.T) {
	sel := Doc2().Find(".even")
	sel2 := Doc2().Find(".five")
	sel3 := sel.NextFilteredUntilSelection(".even", sel2).End()
	AssertEqual(t, sel, sel3)
}

func TestNextFilteredUntilNodes(t *testing.T) {
	sel := Doc2().Find(".even")
	sel2 := Doc2().Find(".four")
	sel3 := sel.NextFilteredUntilNodes(".odd", sel2.Nodes...)
	AssertLength(t, sel3.Nodes, 4)
	AssertSelectionIs(t, sel3, "#n2", "#n6", "#nf2", "#nf6")

	selc := sel.NextFilteredUntilNodesCompiled(cascadia.MustCompile(".odd"), sel2.Nodes...)
	AssertSelEqual(t, sel3, selc)
}

func TestNextFilteredUntilNodesRollback(t *testing.T) {
	sel := Doc2().Find(".even")
	sel2 := Doc2().Find(".four")
	sel3 := sel.NextFilteredUntilNodes(".odd", sel2.Nodes...).End()
	AssertEqual(t, sel, sel3)
}

func TestPrevFilteredUntil(t *testing.T) {
	sel := Doc2().Find(".five").PrevFilteredUntil(".odd", ".one")
	AssertLength(t, sel.Nodes, 4)
	AssertSelectionIs(t, sel, "#n4", "#n2", "#nf4", "#nf2")

	selc := Doc2().Find(".five").PrevFilteredUntilCompiled(
		cascadia.MustCompile(".odd"),
		cascadia.MustCompile(".one"))
	AssertSelEqual(t, sel, selc)
}

func TestPrevFilteredUntilRollback(t *testing.T) {
	sel := Doc2().Find(".four")
	sel2 := sel.PrevFilteredUntil(".odd", ".one").End()
	AssertEqual(t, sel, sel2)
}

func TestPrevFilteredUntilSelection(t *testing.T) {
	sel := Doc2().Find(".odd")
	sel2 := Doc2().Find(".two")
	sel3 := sel.PrevFilteredUntilSelection(".odd", sel2)
	AssertLength(t, sel3.Nodes, 2)
	AssertSelectionIs(t, sel3, "#n4", "#nf4")

	selc := sel.PrevFilteredUntilSelectionCompiled(cascadia.MustCompile(".odd"), sel2)
	AssertSelEqual(t, sel3, selc)
}

func TestPrevFilteredUntilSelectionRollback(t *testing.T) {
	sel := Doc2().Find(".even")
	sel2 := Doc2().Find(".five")
	sel3 := sel.PrevFilteredUntilSelection(".even", sel2).End()
	AssertEqual(t, sel, sel3)
}

func TestPrevFilteredUntilNodes(t *testing.T) {
	sel := Doc2().Find(".even")
	sel2 := Doc2().Find(".four")
	sel = sel.PrevFilteredUntilNodes(".odd", sel2.Nodes...)
	AssertLength(t, sel.Nodes, 2)
	AssertSelectionIs(t, sel, "#n2", "#nf2")
}

func TestPrevFilteredUntilNodesRollback(t *testing.T) {
	sel := Doc2().Find(".even")
	sel2 := Doc2().Find(".four")
	sel3 := sel.PrevFilteredUntilNodes(".odd", sel2.Nodes...).End()
	AssertEqual(t, sel, sel3)
}

func TestClosestItself(t *testing.T) {
	sel := Doc2().Find(".three")
	sel2 := sel.Closest(".row")
	AssertLength(t, sel2.Nodes, sel.Length())
	AssertSelectionIs(t, sel2, "#n3", "#nf3")

	selc := sel.ClosestCompiled(cascadia.MustCompile(".row"))
	AssertSelEqual(t, sel2, selc)
}

func TestClosestNoDupes(t *testing.T) {
	sel := Doc().Find(".span12")
	sel2 := sel.Closest(".pvk-content")
	AssertLength(t, sel2.Nodes, 1)
	AssertClass(t, sel2, "pvk-content")

	selc := sel.ClosestCompiled(cascadia.MustCompile(".pvk-content"))
	AssertSelEqual(t, sel2, selc)
}

func TestClosestNone(t *testing.T) {
	sel := Doc().Find("h4")
	sel2 := sel.Closest("a")
	AssertLength(t, sel2.Nodes, 0)

	selc := sel.ClosestCompiled(cascadia.MustCompile("a"))
	AssertSelEqual(t, sel2, selc)
}

func TestClosestMany(t *testing.T) {
	sel := Doc().Find(".container-fluid")
	sel2 := sel.Closest(".pvk-content")
	AssertLength(t, sel2.Nodes, 2)
	AssertSelectionIs(t, sel2, "#pc1", "#pc2")

	selc := sel.ClosestCompiled(cascadia.MustCompile(".pvk-content"))
	AssertSelEqual(t, sel2, selc)
}

func TestClosestRollback(t *testing.T) {
	sel := Doc().Find(".container-fluid")
	sel2 := sel.Closest(".pvk-content").End()
	AssertEqual(t, sel, sel2)
}

func TestClosestSelectionItself(t *testing.T) {
	sel := Doc2().Find(".three")
	sel2 := sel.ClosestSelection(Doc2().Find(".row"))
	AssertLength(t, sel2.Nodes, sel.Length())
}

func TestClosestSelectionNoDupes(t *testing.T) {
	sel := Doc().Find(".span12")
	sel2 := sel.ClosestSelection(Doc().Find(".pvk-content"))
	AssertLength(t, sel2.Nodes, 1)
	AssertClass(t, sel2, "pvk-content")
}

func TestClosestSelectionNone(t *testing.T) {
	sel := Doc().Find("h4")
	sel2 := sel.ClosestSelection(Doc().Find("a"))
	AssertLength(t, sel2.Nodes, 0)
}

func TestClosestSelectionMany(t *testing.T) {
	sel := Doc().Find(".container-fluid")
	sel2 := sel.ClosestSelection(Doc().Find(".pvk-content"))
	AssertLength(t, sel2.Nodes, 2)
	AssertSelectionIs(t, sel2, "#pc1", "#pc2")
}

func TestClosestSelectionRollback(t *testing.T) {
	sel := Doc().Find(".container-fluid")
	sel2 := sel.ClosestSelection(Doc().Find(".pvk-content")).End()
	AssertEqual(t, sel, sel2)
}

func TestClosestNodesItself(t *testing.T) {
	sel := Doc2().Find(".three")
	sel2 := sel.ClosestNodes(Doc2().Find(".row").Nodes...)
	AssertLength(t, sel2.Nodes, sel.Length())
}

func TestClosestNodesNoDupes(t *testing.T) {
	sel := Doc().Find(".span12")
	sel2 := sel.ClosestNodes(Doc().Find(".pvk-content").Nodes...)
	AssertLength(t, sel2.Nodes, 1)
	AssertClass(t, sel2, "pvk-content")
}

func TestClosestNodesNone(t *testing.T) {
	sel := Doc().Find("h4")
	sel2 := sel.ClosestNodes(Doc().Find("a").Nodes...)
	AssertLength(t, sel2.Nodes, 0)
}

func TestClosestNodesMany(t *testing.T) {
	sel := Doc().Find(".container-fluid")
	sel2 := sel.ClosestNodes(Doc().Find(".pvk-content").Nodes...)
	AssertLength(t, sel2.Nodes, 2)
	AssertSelectionIs(t, sel2, "#pc1", "#pc2")
}

func TestClosestNodesRollback(t *testing.T) {
	sel := Doc().Find(".container-fluid")
	sel2 := sel.ClosestNodes(Doc().Find(".pvk-content").Nodes...).End()
	AssertEqual(t, sel, sel2)
}

func TestIssue26(t *testing.T) {
	img1 := `<img src="assets/images/gallery/thumb-1.jpg" alt="150x150" />`
	img2 := `<img alt="150x150" src="assets/images/gallery/thumb-1.jpg" />`
	cases := []struct {
		s string
		l int
	}{
		{s: img1 + img2, l: 2},
		{s: img1, l: 1},
		{s: img2, l: 1},
	}
	for _, c := range cases {
		doc, err := NewDocumentFromReader(strings.NewReader(c.s))
		if err != nil {
			t.Fatal(err)
		}
		sel := doc.Find("img[src]")
		AssertLength(t, sel.Nodes, c.l)
	}
}
