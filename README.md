## LeetCode Go Practice

Personal Go solutions for selected LeetCode problems, organized by topic.

### Current Problems

- **array**
  - `p1461.go`: Check If a String Contains All Binary Codes of Size K

- **linkedlist**
  - `p002_add_two_numbers.go`: Add Two Numbers (linked list representation)

- **string**
  - `p003_longest_substring_without_dup.go`: Longest Substring Without Repeating Characters
  - `p2981.go`: Maximum Length of a String With Repeated Characters

More problems will be added over time, grouped by data structure / topic and file name.

### Project Structure

```text
leetcode-go/
├── array/            # array-related problems
│   └── p1461.go
├── linkedlist/       # linked list problems
│   └── p002_add_two_numbers.go
├── string/           # string problems
│   ├── p003_longest_substring_without_dup.go
│   └── p2981.go
└── README.md
```

### Build & Test

From the project root:

```bash
go build ./...
go test ./... -v
```

If you add test files for a specific problem (for example `array/p1461_test.go`), you can run tests only for that package:

```bash
cd array
go test . -v
```

or from the project root:

```bash
go test ./array -v
```

### Conventions for New Problems

- **File naming**: `p<id>_<short_english_name>.go` or simply `p<id>.go`
  - e.g. `p001_two_sum.go`, `p217_contains_duplicate.go`
- **Directory grouping**: by data structure / topic
  - Array: `array/`
  - String: `string/`
  - Linked list: `linkedlist/`
  - Can be extended later: `tree/`, `dp/`, etc.

Follow these conventions when adding new solutions to keep the repository organized and easy to review.
