# Sock Merchant

## Problem

John works at a clothing store. He has a large pile of socks that he must pair by color for sale. Given an array of integers representing the color of each sock, determine how many pairs of socks with matching colors there are.

## Step

1. วนลูปเอาถุงเท้าข้างแรกมาตั้งเพื่อจะเทียบอีกถุงเท้าอีกข้าง
2. วนเอาถุงเท้าอีกข้างมาเพียบเพื่อเอาตัวเลขที่เหมือนกันมาจับคู่
3. ถ้าจับคู่กันได้แล้ว ให้นับ 1 และ set ตัวเลขให้เป็น 0
4. วนแบบนี้จนไม่มีถุงเท้าที่เหมือน
5. return ว่ามีทั้งหมดกี่คู่
