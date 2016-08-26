package main

var fabs = [100]uint64{1, 1}
var fabs_calcued uint32 = 2

func fab(seed uint32) uint64 {
	if seed > uint32(len(fabs)) {
		panic(fab)
	} // 这里的else省略了很自然

	if fabs[seed-1] != 0 {
		return fabs[seed-1]
	} else if seed == fabs_calcued+1 {
		// TODO: 越界检查，求和是否已经超过uint64最大值
		fabs[fabs_calcued] = fabs[fabs_calcued-1] + fabs[fabs_calcued-2]
		fabs_calcued++
		return fabs[seed-1]
	} else { // 这个else省略了会让代码难读，一般只在针对异常处理的单个else进行省略
		for i := fabs_calcued + 1; i < seed; i++ {
			fab(i)
		}
		return fab(seed)
	}
}
