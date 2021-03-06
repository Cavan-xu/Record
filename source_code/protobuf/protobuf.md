# protobuf 源码阅读
protobuf 的核心编码原理，包括 Varint 编码、ZigZag编码及 protobuf 特有的 Message Structure 编码结构

##### Varint 编码：

> protobuf 编码主要依赖于 Varint 编码

Varint 是一种紧凑的表示数字的方法。它用一个或多个字节来表示一个数字，值越小的数字使用越少的字节数。这能减少用来表示数字的字节数。Varint 中的每个字节（最后一个字节除外）都设置了最高有效位（msb），这一位表示是否还会有更多字节出现。每个字节的低 7 位用于以 7 位组的形式存储数字的二进制补码表示，最低有效组首位。

> 最高位为1代表后面7位仍然表示数字，否则为0，后面7位用原码补齐。

如果用不到 1 个字节，那么最高有效位设为 0 ，如下面这个例子，1 用一个字节就可以表示，所以 msb 为 0.

| `1 ` | `00000001 ` |
| ---- | ----------- |

```shell
二进制表示：00000001
分割七位：0 0000001
没有其他字节出现，高位补0，即 00000001
```

如果需要多个字节表示，msb 就应该设置为 1 。例如 300，如果用 Varint 表示的话：

| `300` | `10101100 00000010 ` |
| ----- | -------------------- |

```shell
二进制表示：00000001 00101100
分割七位：00 0000010 0101100
转化为小端序：0101100 0000010	// 低地址->高地址	小端序：低地址放字节低位，所以从 0000010 0101100 转化为 0101100 0000010
varint编码：10101100 00000010
```

> [字节序](https://baike.baidu.com/item/%E5%AD%97%E8%8A%82%E5%BA%8F/1457160?fr=aladdin)

![](https://raw.githubusercontent.com/Cavan-xu/Images/master/Record/protobuf/protobuf-1.png)

编码方式：

1. 将被编码数转化为二进制表示
2. 从低位到高位按 7 位一组划分
3. 将大端序转化为小端序，即以分组为单位进行顺序交换，protobuf 默认使用小端序
4. 给每组加上最高有效位(最后一个字节高位补0，其余各字节高位补1)组成编码后的数据
5. 最后转成 10 进制

解码过程：

就是将字节依次取出，去掉最高有效位，因为是小端排序所以先解码的字节要放在低位，之后解码出来的二进制位继续放在之前已经解码出来的二进制的高位最后转换为10进制数完成varint编码的解码过程。

##### Varint 编码的缺点：

负数需要10个字节显示（因为计算机定义负数的符号位为数字的最高位）。（ps: 64 / 7 = 10）

> 具体是先将负数是转成了 long 类型，再进行 varint 编码，这就是占用 10个 字节的原因了。

![](https://raw.githubusercontent.com/Cavan-xu/Images/master/Record/protobuf/protobuf-2.png)

**protobuf 采取的解决方式：使用 sint32/sint64 类型表示负数，通过先采用 Zigzag 编码，将正数、负数和0都映射到无符号数，最后再采用 varint 编码**。

##### ZigZag 编码：

