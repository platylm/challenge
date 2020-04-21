# Jumping on the Clouds

## Problem

Emma is playing a new mobile game that starts with consecutively numbered clouds. Some of the clouds are thunderheads and others are cumulus. She can jump on any cumulus cloud having a number that is equal to the number of the current cloud plus  or . She must avoid the thunderheads. Determine the minimum number of jumps it will take Emma to jump from her starting postion to the last cloud. It is always possible to win the game.

## Step
1. ถ้าเป็น 0 ให้นับ 1 ถ้าเป็น 1 ไม่นับ
2. ให้เช็ค array + [2] ก่อน ถ้าเป็น 0 ให้นับ 1 ถ้า array + [2] แล้วเป็น 1 ให้มาเช็ค array + [1]