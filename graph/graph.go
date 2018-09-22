package goraph

//Vertex : struct of vertex that have id and weight
type Vertex struct {
	ID     int
	Weight float32
}

//Edge : struct of edge. Maybe this struct don't need...
type Edge struct {
	From   Vertex
	To     Vertex
	Weight float32
}

//Graph : グラフの構造体。privateにすべき？
//グラフは辺集合を持つべき？隣接行列かLinkedList持ってれば持たなくてもいいが、隣接行列とLinkeListどっちを持てばいいかはグラフの疎密具合に依存する。。。
//辺集合を持っておいて後でユーザー側でBuildAdjMatrix or BuildLinkedListみたいなのをしたほうがいい気もする
//隣接行列なら辺の問い合わせは定数。隣接行列持った方がいいかな。
//隣接行列だけもつ。頂点の追加は隣接行列に行と列を足す。辺の追加は該当要素に1を立てる。とすれば良いのか？？？？
//いやーMST求める時とかはEdge Set持ってた方が楽だったりするんだよな。。。
type Graph struct {
}
