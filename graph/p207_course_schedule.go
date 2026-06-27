package graph

// p207 课程表
// 你需要选修 numCourses 门课程，标记为 0 到 numCourses-1。
// 给你一个数组 prerequisites，其中 prerequisites[i] = [ai, bi] 表示：
// 如果你想要选修 ai 课程，你必须先选修 bi 课程。
//
// 例如，你想学习课程 0，你需要先完成课程 1，用 [0, 1] 表示。
// 请你判断是否可能完成所有课程的学习（即是否存在拓扑排序顺序）。
//
// 示例 1：
// 输入：numCourses = 2, prerequisites = [[1, 0]]
// 输出：true
// 解释：可以按 0 -> 1 的顺序完成课程。
//
// 示例 2：
// 输入：numCourses = 2, prerequisites = [[1, 0], [0, 1]]
// 输出：false
// 解释：存在循环依赖，无法完成。
//
// 提示：
// - 1 <= numCourses <= 2000
// - 0 <= prerequisites.length <= 5000
// - prerequisites[i].length == 2
// - 0 <= ai, bi < numCourses
// - 所有先修条件对都不同

func canFinish(numCourses int, prerequisites [][]int) bool {

}
