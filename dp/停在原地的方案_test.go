package dp

import (
	"fmt"
	"math"
	"testing"
)

func Test停在原地的方案(t *testing.T) {
	mod := int64(1e9 + 7)
	steps, stepLen := 3, 2
	max := int(math.Min(float64(steps/2), float64(stepLen-1)))
	f := make([][]int64, steps+1)
	for i := 0; i < steps+1; i++ {
		f[i] = make([]int64, max+1)
	}
	f[steps][0] = 1
	for i := steps - 1; i >= 0; i-- {
		for j := 0; j <= max; j++ {
			f[i][j] = (f[i][j] + f[i+1][j]) % mod
			if j-1 >= 0 {
				f[i][j] = (f[i][j] + f[i+1][j-1]) % mod
			}
			if j+1 <= max {
				f[i][j] = (f[i][j] + f[i+1][j+1]) % mod
			}
		}
	}
	fmt.Println(f[0][0])
}

// 题解：https://juejin.cn/post/6963870119235682311?utm_source=gold_browser_extension
// class Solution {
//     int mod = (int)1e9+7;
//     public int numWays(int steps, int len) {
//         int max = Math.min(steps / 2, len - 1);
//         int[][] f = new int[steps + 1][max + 1];
//         f[steps][0] = 1;
//         for (int i = steps - 1; i >= 0; i--) {
//             for (int j = 0; j <= max; j++) {
//                 f[i][j] = (f[i][j] + f[i + 1][j]) % mod;
//                 if (j - 1 >= 0) f[i][j] = (f[i][j] + f[i + 1][j - 1]) % mod;
//                 if (j + 1 <= max) f[i][j] = (f[i][j] + f[i + 1][j + 1]) % mod;
//             }
//         }
//         return f[0][0];
//     }
// }
