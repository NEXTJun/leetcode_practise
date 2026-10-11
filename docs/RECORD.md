# LeetCode 練習紀錄

## 使用方式

- `狀態`：Not started / Solving / Solved / Review / Revisit
- `難度`：Easy / Medium / Hard
- `核心想法`：用一句話寫出本題真正考點。
- `卡住點`：記錄錯誤假設、邊界案例、語法或資料結構問題。
- `下次複習`：建議間隔為 1 天、3 天、7 天、14 天。

## 總表

| 日期 | 題號 | 題目 | 難度 | 狀態 | 核心想法 | 時間複雜度 | 空間複雜度 | 卡住點 | 下次複習 |
|---|---:|---|---|---|---|---|---|---|---|
| 2026-05-03 | 1 | Two Sum | Easy | Solved | 用 hash map 記錄已看過的數字，查 complement | O(n) | O(n) | 建立 map 前後順序會影響是否用到同一個元素 | 待定 |
| 2026-05-03 | 20 | Valid Parentheses | Easy | Solved | 用 stack 檢查括號是否符合後進先出 | O(n) | O(n) | 不需要額外 res；錯誤時直接 return false，最後檢查 stack 是否清空 | 待定 |
| 2026-09-29 | 121 | Best Time to Buy and Sell Stock | Easy | Solved | 逐日當賣出日，維護先前最低買價與歷史最佳利潤 | O(n) | O(1) | 本輪由 O(n²) 暴力法轉為單次走訪；已獨立實作並口述先買後賣與複雜度 | 2026-10-02 |
| 2026-10-04 | 198 | House Robber | Medium | Solved | 前綴最佳分拿與不拿，DP 比較 F[i-1] 與 F[i-2]+nums[i] | O(n) | O(n) | 本輪修正 i=1 落入通用轉移造成 F[-1]；通用正確性口述、邊界口述與 O(1) 空間版待複習 | 2026-10-07 |
| 2026-05-04 | 217 | Contains Duplicate | Easy | Solved | 用 hash set 記錄已看過的數字，遇到重複就提早回傳 | O(n) | O(n) | 解法正確；命名可從 data 改成 seen 更清楚 | 待定 |
| 2026-05-04 | 242 | Valid Anagram | Easy | Solved | 排序法直觀；小寫英文字母限制下可用 [26]int 計數法優化 | 排序 O(n log n)；計數 O(n) | 排序 O(n)；計數 O(1) | 兩種方法都要會：排序法好口述，計數法是本題限制下複雜度最優 | 待定 |
| 2026-09-28 | 704 | Binary Search | Easy | Solved | 閉區間二分搜尋依中點大小排除半邊，直到找到或區間為空 | O(log n) | O(1) | 本輪修正左右移動方向、初始 right 與漏檢最後一格；已能口述原因及複雜度 | 2026-10-05 |
| 2026-05-05 | 125 | Valid Palindrome | Easy | Solved | 雙指標跳過非英數字元，轉小寫後比較 | O(n) | O(1) | 指標移動時機要小心；跳過符號後若左右相遇應視為完成 | 待定 |
| 2026-09-28 | 53 | Maximum Subarray | Medium | Solved | 以目前元素結尾的最佳和比較接續或重開，另維護全域最大 | O(n) | O(1) | 本輪修正 current/best 混淆及全負數初始化；口述 current 的精確定義可再練 | 2026-10-05 |
| 2026-05-07 | 35 | Search Insert Position | Easy | Solved | Binary Search 找不到時回傳 left 作為插入位置 | O(log n) | O(1) | 二分搜尋邊界用 index 更新，不要把 nums[mid] 當成 index | 待定 |
| 2026-05-07 | 278 | First Bad Version | Easy | Solved | 單調 good/bad 狀態用 binary search 找第一個 bad | O(log n) | O(1) | `right = mid` 要搭配 `left < right`；good 時要 `left = mid + 1` | 待定 |
| 2026-05-07 | 383 | Ransom Note | Easy | Solved | 統計 magazine 字母庫存，再用 ransomNote 扣庫存 | O(n + m) | map O(k)；陣列 O(1) | map 可通用；小寫限制下 [26]int 空間更好，可扣到負數時提早回傳 | 待定 |
| 2026-09-28 | 167 | Two Sum II - Input Array Is Sorted | Medium | Solved | 已排序陣列用 left/right 雙指標依 sum 大小收縮 | O(n) | O(1) | 本輪修正索引與 1-based 回傳；排除一側的嚴謹論證待複習 | 2026-10-05 |
| 2026-05-07 | 226 | Invert Binary Tree | Easy | Solved | 對每個節點交換左右子樹，再遞迴處理子樹 | O(n) | O(h) | recursion 要有 root == nil 停止條件；可用多重賦值交換左右 | 待定 |
| 2026-05-07 | 100 | Same Tree | Easy | Solved | 遞迴比較兩棵樹對應節點的結構和值 | O(n) | O(h) | 遞迴要比較 p.Left 對 q.Left、p.Right 對 q.Right，不是同樹左右互比 | 待定 |
| 2026-05-10 | 101 | Symmetric Tree | Easy | Solved | 遞迴交叉比較左右子樹是否互為鏡像 | O(n) | O(h) | 對稱比較是 left.Left 對 right.Right、left.Right 對 right.Left | 待定 |
| 2026-05-10 | 104 | Maximum Depth of Binary Tree | Easy | Solved | 遞迴取左右子樹較大深度再加上目前節點 | O(n) | O(h) | 空樹深度是 0；目前節點本身要 +1 | 待定 |
| 2026-05-10 | 110 | Balanced Binary Tree | Easy | Solved | bottom-up 遞迴回傳高度，不平衡時用 -1 往上傳 | O(n) | O(h) | 每個節點都要檢查左右高度差；用 -1 避免重複計算高度 | 待定 |
| 2026-09-29 | 70 | Climbing Stairs | Easy | Solved | 依最後一步由 n-1 或 n-2 抵達，memo 避免重算走法數 | O(n) | memo O(n)；rolling O(1) | 本輪修正 helper 基底停止與共用 map；rolling O(1) 尚未獨立實作 | 2026-10-02 |
| 2026-05-10 | 733 | Flood Fill | Easy | Solved | DFS 從起點往四方向擴散，只填原顏色相同的連通區塊 | O(m*n) | O(m*n) | 停止條件是顏色不同或越界；originalColor == color 要直接返回 | 待定 |
| 2026-05-16 | 542 | 01 Matrix | Medium | Solved | multi-source BFS 從所有 0 同時擴散，第一次到達即最短距離 | O(m*n) | O(m*n) | 用 dist == -1 防重複；`queue = queue[1:]` 不能寫成 `:=` | 待定 |
| 2026-05-16 | 235 | Lowest Common Ancestor of a Binary Search Tree | Medium | Solved | 利用 BST 大小關係，p/q 同側就往同側走，否則 root 是 LCA | O(h) | recursion O(h)；iteration O(1) | LCA 可是 p 或 q 本身；BST 節點值不重複，可用 Val 比大小 | 待定 |
| 2026-09-28 | 98 | Validate Binary Search Tree | Medium | Solved | DFS 傳遞嚴格上下界，檢查每個節點是否符合所有祖先限制 | O(n) | O(h) | 本輪修正漏檢右子樹及界線傳遞；遞迴堆疊複雜度仍待複習 | 2026-09-29 |
| 2026-05-16 | 112 | Path Sum | Easy | Solved | DFS 扣剩餘 target，只有走到 leaf 才判斷是否等於目標 | O(n) | O(h) | 中途加總等於 target 不算成功；必須是 root-to-leaf 完整路徑 | 待定 |
| 2026-05-16 | 105 | Construct Binary Tree from Preorder and Inorder Traversal | Medium | Solved | preorder 第一個找 root，inorder 找 root 位置切左右子樹 | O(n^2) | O(n) | slice 版好寫；map 可把找 root 位置優化到 O(1)，整體 O(n) | 待定 |
| 2026-05-20 | 102 | Binary Tree Level Order Traversal | Medium | Solved | BFS 用 queue 按層處理；DFS 版用 depth 放入對應層 | O(n) | BFS O(n)；DFS O(h) 不含結果 | BFS 要固定 levelSize；小孩放 queue，不放目前 level | 待定 |
| 2026-06-07 | 199 | Binary Tree Right Side View | Medium | Solved | BFS 按層處理，每層最後一個節點就是右視圖 | O(n) | O(n) | 不是只走 .Right；左右小孩都要進 queue，每層固定 levelSize | 待定 |
| 2026-09-28 | 200 | Number of Islands | Medium | Solved | 掃描未處理陸地後以 DFS 標記整座島，避免重複計數 | O(m*n) | O(m*n) | 本輪修正 map key、byte 字元、越界與標記時機；已能口述計數理由和複雜度 | 2026-10-05 |

## 每題筆記模板

### 題號：題目名稱

- 日期：
- 難度：
- 題型：
- 第一次想法：
- 最佳解法：
- 正確性理由：
- 時間複雜度：
- 空間複雜度：
- 邊界案例：
- 錯誤紀錄：
- 面試口述版本：
- 下次複習：
