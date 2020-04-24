# Repeated String
## Problem
Lilah has a string, s, of lowercase English letters that she repeated infinitely many times.

Given an integer, n, find and print the number of letter a's in the first n letters of Lilah's infinite string.

For example, if the string s = 'abcac' and n = 10 , the substring we consider is abcacabcac, the first 10 characters of her infinite string. There are 4 occurrences of a in the substring.

## Step
1. นับตัว a ที่อยู่ในชุด string ก่อนว่ามีกี่ตัว
2. เอาจำนวน n มาหารชุด string ว่าจะต้อง repeat กี่ครั้ง
3. จากนั้นเอาจำนวน n % ชุด string เอาเศษว่าจะต้อง repeat อีกกี่ตัว
4. รวมจำนวน a จากจำนวน repeat และเศษที่เหลือที่มีตัว a
