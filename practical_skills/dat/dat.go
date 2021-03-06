package dat

/*
	double array trie：双数组前缀树
	背景：项目中玩家昵称、聊天信息涉及到敏感字的加敏。
	最开始的实现：使用普通的前缀树，因为普通前缀树如果想要快速访问，必须使用 map 来做空间换时间，由于敏感字的数量比较大，
				导致项目占用内存增长了大概 500mb，加载前缀树之前只有 40mb，这是不能接受的
	优化：引入双数组前缀树，因为敏感字只需要在程序启动的时候 load 到内存里，加入之后并不需要增减，因此十分适合双数组前缀树，并且内存的占用并不高
*/
