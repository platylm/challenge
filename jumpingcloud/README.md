# Jumping on the Clouds

## Problem

Emma is playing a new mobile game that starts with consecutively numbered clouds. Some of the clouds are thunderheads and others are cumulus. She can jump on any cumulus cloud having a number that is equal to the number of the current cloud plus  or . She must avoid the thunderheads. Determine the minimum number of jumps it will take Emma to jump from her starting postion to the last cloud. It is always possible to win the game.

## Step

1. เอาชุดตัวเลข cloud มา for loop เพื่อจะเช็คตัวเลขแต่ละตำแหน่ง
2. ถ้าเป็น 0 ให้นับ 1 ถ้าเป็น 1 ไม่นับและให้ข้ามตัวถัดไป
3. ให้เช็ค array + [2] ก่อน ถ้าเป็น 0 ให้นับ 1 และขยับมาตำแหน่งที่ array + [2] แต่ถ้า array + [2] แล้วเป็น 1 ไม่นับและให้อยู่ในตำแหน่งเดิม
4. ถ้า array + [2] ไม่ใช่ 0 ให้มาเช็ค array + [1] ถ้าเป็น 0 ให้นับ 1 และขยับมาตำแหน่งที่ array + [1]
5. ทำซ้ำจนกว่าจะครบทุกตัวในชุดตัวเลข cloud
6. return จำนวนเลข 0 ที่สามารถกระโดดได้ว่ามีกี่ครั้ง
