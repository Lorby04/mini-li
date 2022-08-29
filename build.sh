rm -rf main `find . -name *.o`
gccgo -o github.com/emirpasic/gods/utils.o -c github.com/emirpasic/gods/utils/utils.go github.com/emirpasic/gods/utils/comparator.go github.com/emirpasic/gods/utils/sort.go
gccgo -o github.com/emirpasic/gods/containers.o -c github.com/emirpasic/gods/containers/containers.go github.com/emirpasic/gods/containers/enumerable.go github.com/emirpasic/gods/containers/iterator.go github.com/emirpasic/gods/containers/serialization.go
gccgo -o github.com/emirpasic/gods/maps.o -c github.com/emirpasic/gods/maps/maps.go
gccgo -o github.com/emirpasic/gods/trees.o -c github.com/emirpasic/gods/trees/trees.go
gccgo -o github.com/emirpasic/gods/trees/redblacktree.o -c github.com/emirpasic/gods/trees/redblacktree/redblacktree.go github.com/emirpasic/gods/trees/redblacktree/iterator.go github.com/emirpasic/gods/trees/redblacktree/serialization.go
gccgo -o github.com/emirpasic/gods/maps/treemap.o -c github.com/emirpasic/gods/maps/treemap/treemap.go github.com/emirpasic/gods/maps/treemap/enumerable.go github.com/emirpasic/gods/maps/treemap/iterator.go github.com/emirpasic/gods/maps/treemap/serialization.go
gccgo -o mini-li/lis/target.o -c mini-li/lis/target/target.go
gccgo -o mini-li/lis/lis_map.o -c mini-li/lis/lis_map/lis.go mini-li/lis/lis_map/litest.go
gccgo -o mini-li/lis/lis_tree.o -c mini-li/lis/lis_tree/lis.go mini-li/lis/lis_tree/litest.go
gccgo -o main -c mini-li/main/mini-li.go
