## 第三课 神奇的内置数据结构

* 封装一个数据结构MyMap. 实现并发安全的Load. Store, Delete, LoadAndDelete, LoadOrStore几个API(禁止使用sync.Map), 不用考虑性能.
* 编写Benchmark, 比较MyMap与sync.Map的同名函数性能差异. 输出相应的性能报告(注意应该使用RunParallel), 将性能比较结果输出为Markdown文件
* 使用channel实现一个trylock
* 修复deadlock.go中的死锁
* 实现一个MyContext, 可以在父节点上查到子节点WithValue中的值

