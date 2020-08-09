# numbers-of-good-pairs

## 題目解讀：
### 題目來源:

[number-of-good-pairs](https://leetcode.com/problems/number-of-good-pairs/)

### 原文:
Given an array of integers nums.

A pair (i,j) is called good if nums[i] == nums[j] and i < j.

Return the number of good pairs.

### 解讀：
給定一個正整數陣列 nums

定義 

(i, j)為good pair

如果i,j 符合

1 nums[i] === nums[j]
2 i < j

找出在 nums 中的所有good pair


## 初步解法:
### 初步觀察:

假設用一個 map儲存 每個值value出現的次數 e.g map[int]int

每個累計值 index後面的值 每多出一個相同值

則pair新增的值為上個 map[value]

因為後面的index剛好都可以跟前面出現過得值產生一個good pair

舉例來說：

[1, 1, 1, 1]
當i = 0 map[1] =1 
count = 0
i = 1 map[1] = 2
count = count + 1(上一個 map[1]) = 1
i = 2 map[1] = 3
count = count + 2(上一個 map[1]) = 3
i = 3 map[1] = 4
count = count + 3(上一個 map[1]) = 6
### 初步設計:
given a integer array nums

set a integer result = 0

set valueMap = map[int]int

loop index i = 0 to i < length of nums:

check valueMap[nums[i]] exists or not

if valueMap[nums[i]] exists set result += valueMap[nums[i]]

valueMap[nums[i]]+=1

return result

## 遇到的困難
### golang map宣告不熟
由於不常使用golang map

因此花了些時間查訊golang map用法
### 題目上理解的問題
因為英文不是筆者母語

所以在題意解讀上 容易被英文用詞解讀給搞模糊

### pseudo code撰寫

一開始不習慣把pseudo code寫下來

因此 不太容易把自己的code做解析

### golang table driven test不熟
對於table driven test還不太熟析

所以對於寫test還是耗費不少時間
## 我的github source code
```golang
package num_identical

func numIdenticalPairs(nums []int) int {
	ret := 0
	valueMap := make(map[int]int)
	for _, val := range nums {
		if v, ok := valueMap[val]; ok == true {
			ret += v
		}
		valueMap[val] += 1
	}
	return ret
}

```
## 測資的撰寫
```golang
package num_identical

import "testing"

func Test_numIdenticalPairs(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example1",
			args: args{
				nums: []int{1, 2, 3, 1, 1, 3},
			},
			want: 4,
		},
		{
			name: "Example2",
			args: args{
				nums: []int{1, 1, 1, 1},
			},
			want: 6,
		},
		{
			name: "Example3",
			args: args{
				nums: []int{1, 2, 3},
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numIdenticalPairs(tt.args.nums); got != tt.want {
				t.Errorf("numIdenticalPairs() = %v, want %v", got, tt.want)
			}
		})
	}
}

```


## 參考文章

[golang test](https://ithelp.ithome.com.tw/articles/10204692)
[golang map宣告](https://michaelchen.tech/golang-programming/map/)