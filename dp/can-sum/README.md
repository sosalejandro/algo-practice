# Graphical Explanation

## Initialization

- A global map `memo` is created to store results of subproblems.

## Function Call

- `CanSum(target, arr)` is called with a target sum and an array of integers.

## Base Cases

- If `target` is `0`, return `true` (we have found a combination that sums to the target).
- If `target` is less than `0`, return `false` (we cannot have a negative target).

## Recursive Case

- Iterate through each element in the array.
- Calculate the remainder `r` by subtracting the current element from the target.
- Recursively call `CanSum(r, arr)` and store the result in `memo[target]`.
- If `memo[target]` is `true`, return `true`.

## Return `false`

- If no combination sums to the target, return `false`.

## Diagram

```plaintext
Initial Call: CanSum(target, arr)
|
|-- Check if target == 0 -> Return true
|-- Check if target < 0 -> Return false
|
|-- For each element in arr:
|   |
|   |-- Calculate r = target - arr[i]
|   |-- Check memo[target] = CanSum(r, arr)
|   |-- If memo[target] is true, return true
|
|-- If no combination found, return false
```
## Example

Let's say `target = 7` and `arr = [5, 3, 4, 7]`.

1. **Initial Call**: `CanSum(7, [5, 3, 4, 7])`
   - `target` is not `0` and not less than `0`.

2. **Iteration**:
   - For `arr[0] = 5`:
     - `r = 7 - 5 = 2`
     - `CanSum(2, [5, 3, 4, 7])`
       - `target` is not `0` and not less than `0`.
       - For `arr[0] = 5`:
         - `r = 2 - 5 = -3`
         - `CanSum(-3, [5, 3, 4, 7])` -> `false`
       - For `arr[1] = 3`:
         - `r = 2 - 3 = -1`
         - `CanSum(-1, [5, 3, 4, 7])` -> `false`
       - For `arr[2] = 4`:
         - `r = 2 - 4 = -2`
         - `CanSum(-2, [5, 3, 4, 7])` -> `false`
       - For `arr[3] = 7`:
         - `r = 2 - 7 = -5`
         - `CanSum(-5, [5, 3, 4, 7])` -> `false`
     - `memo[2] = false`
   - For `arr[1] = 3`:
     - `r = 7 - 3 = 4`
     - `CanSum(4, [5, 3, 4, 7])`
       - `target` is not `0` and not less than `0`.
       - For `arr[0] = 5`:
         - `r = 4 - 5 = -1`
         - `CanSum(-1, [5, 3, 4, 7])` -> `false`
       - For `arr[1] = 3`:
         - `r = 4 - 3 = 1`
         - `CanSum(1, [5, 3, 4, 7])`
           - `target` is not `0` and not less than `0`.
           - For `arr[0] = 5`:
             - `r = 1 - 5 = -4`
             - `CanSum(-4, [5, 3, 4, 7])` -> `false`
           - For `arr[1] = 3`:
             - `r = 1 - 3 = -2`
             - `CanSum(-2, [5, 3, 4, 7])` -> `false`
           - For `arr[2] = 4`:
             - `r = 1 - 4 = -3`
             - `CanSum(-3, [5, 3, 4, 7])` -> `false`
           - For `arr[3] = 7`:
             - `r = 1 - 7 = -6`
             - `CanSum(-6, [5, 3, 4, 7])` -> `false`
         - `memo[1] = false`
       - For `arr[2] = 4`:
         - `r = 4 - 4 = 0`
         - `CanSum(0, [5, 3, 4, 7])` -> `true`
       - `memo[4] = true`
     - `memo[7] = true`

3. **Result**:
   - `CanSum(7, [5, 3, 4, 7])` -> `true`

The memoization helps avoid redundant calculations by storing results of subproblems.