rm -rf main `find . -name *.o`
gccgo -o github.com/emirpasic/gods/utils.o -c github.com/emirpasic/gods/utils/utils.go github.com/emirpasic/gods/utils/comparator.go github.com/emirpasic/gods/utils/sort.go
gccgo -o github.com/emirpasic/gods/containers.o -c github.com/emirpasic/gods/containers/containers.go github.com/emirpasic/gods/containers/enumerable.go github.com/emirpasic/gods/containers/iterator.go github.com/emirpasic/gods/containers/serialization.go
gccgo -o github.com/emirpasic/gods/maps.o -c github.com/emirpasic/gods/maps/maps.go
gccgo -o github.com/emirpasic/gods/trees.o -c github.com/emirpasic/gods/trees/trees.go
gccgo -o github.com/emirpasic/gods/trees/redblacktree.o -c github.com/emirpasic/gods/trees/redblacktree/redblacktree.go github.com/emirpasic/gods/trees/redblacktree/iterator.go github.com/emirpasic/gods/trees/redblacktree/serialization.go
gccgo -o github.com/emirpasic/gods/maps/treemap.o -c github.com/emirpasic/gods/maps/treemap/treemap.go github.com/emirpasic/gods/maps/treemap/enumerable.go github.com/emirpasic/gods/maps/treemap/iterator.go github.com/emirpasic/gods/maps/treemap/serialization.go
gccgo -o ./lis/target.o -c ./lis/target/target.go
gccgo -o ./lis/lis_map.o -c ./lis/lis_map/lis.go ./lis/lis_map/litest.go
gccgo -o ./lis/lis_tree.o -c ./lis/lis_tree/lis.go ./lis/lis_tree/litest.go
gccgo -o main -c ./main/mini-li.go
