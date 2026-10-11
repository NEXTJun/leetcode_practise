# Recent 10 LeetCode Review

這份是你最近 10 題的快速複習版，只留面試前最該記的骨架。

## 70. Climbing Stairs

- 題意：每次走 1 階或 2 階，問走到第 `n` 階有幾種方法。
- 核心解法：DP / Fibonacci。
- 轉移：`ways[n] = ways[n-1] + ways[n-2]`
- 口訣：到第 `n` 階，只能從 `n-1` 或 `n-2` 走來。
- 易錯點：滾動變數版只要記前兩項；不要在遞迴 helper 裡一直用外層固定 `n`。

## 733. Flood Fill

- 題意：從起點開始，把同色連通區塊全部改成新顏色。
- 核心解法：DFS 四方向擴散。
- 口訣：越界停、顏色不同停、同色就染再往四周走。
- 易錯點：`originalColor == color` 要先直接 return，不然會無限重刷。

## 542. 01 Matrix

- 題意：對每個 `1`，求到最近 `0` 的距離。
- 核心解法：multi-source BFS。
- 作法：先把所有 `0` 丟進 queue，距離設 `0`，再同時往外擴散。
- 口訣：不是從每個 `1` 找 `0`，而是從所有 `0` 一起出發。
- 易錯點：用 `dist == -1` 判斷是否尚未拜訪；`queue = queue[1:]` 不要寫成 `:=`。

## 235. Lowest Common Ancestor of a BST

- 題意：找 BST 中 `p` 和 `q` 的最近共同祖先。
- 核心解法：利用 BST 大小關係。
- 作法：
  - `p`、`q` 都比 root 小，往左。
  - `p`、`q` 都比 root 大，往右。
  - 否則目前 root 就是答案。
- 口訣：同側就往下，分岔點就是 LCA。
- 易錯點：LCA 可以剛好就是 `p` 或 `q` 本身。

## 98. Validate Binary Search Tree

- 題意：判斷一棵樹是不是合法 BST。
- 核心解法：DFS 帶上下界。
- 作法：每個節點都要滿足 `lower < val < upper`。
- 口訣：不能只看父子，要看整條祖先給的範圍。
- 易錯點：BST 不允許重複，所以比較要用嚴格 `<` / `>`；等於邊界也不合法。

## 112. Path Sum

- 題意：是否存在一條 root-to-leaf 路徑，總和等於 `targetSum`。
- 核心解法：DFS 扣剩餘目標值。
- 作法：往下遞迴時傳 `targetSum - node.Val`。
- 口訣：只有走到 leaf 才能判斷成功。
- 易錯點：中途和剛好等於 target 不算，必須是完整 root-to-leaf。

## 105. Construct Binary Tree from Preorder and Inorder Traversal

- 題意：由 preorder 和 inorder 重建二元樹。
- 核心解法：遞迴切左右子樹。
- 作法：
  - `preorder[0]` 是 root。
  - 去 `inorder` 找 root 位置 `i`。
  - 左子樹大小就是 `i`。
  - 用這個大小去切 preorder / inorder。
- 口訣：preorder 負責告訴你 root 是誰，inorder 負責告訴你左右怎麼切。
- 易錯點：slice 邊界最容易錯；直覺版 `O(n^2)`，map 優化可到 `O(n)`。

## 102. Binary Tree Level Order Traversal

- 題意：回傳每一層的節點值。
- 核心解法：BFS queue。
- 作法：每輪先記 `levelSize := len(queue)`，只處理這一層的節點。
- 口訣：這一輪 queue 裡的人，就是同一層。
- 易錯點：小孩要加進 queue，不是直接加進目前 `level`。
- 補充：DFS 也能做，靠 `depth` 放進 `res[depth]`。

## 199. Binary Tree Right Side View

- 題意：從右邊看樹，回傳每一層看到的節點。
- 核心解法：BFS 按層處理。
- 作法：每層處理時，當 `i == levelSize-1`，把該節點加入答案。
- 口訣：右視圖 = 每一層最後處理到的節點。
- 易錯點：不是一路走 `.Right`；左右小孩都要正常進 queue。

## 200. Number of Islands

- 題意：grid 裡有幾座島。
- 核心解法：掃 grid + DFS 淹島。
- 作法：遇到 `'1'` 就 `count++`，再 DFS 把整座島標記掉。
- 口訣：找到新島，就把整座島淹掉。
- 易錯點：grid 是 `[][]byte`，要用 `'1'` / `'0'`，不是 `1` / `0`。

## 總結口訣

- 問子樹、合法性、路徑：先想 DFS。
- 問每層、最短距離、一圈一圈擴散：先想 BFS。
- 問 grid 連通塊：先想 DFS/BFS + visited 標記。
- BFS 常見關鍵：先固定 `levelSize`。
- Grid 常見關鍵：越界、visited、四方向。
