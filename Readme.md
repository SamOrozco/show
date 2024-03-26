# Show
A command line utility for managing commonly used links and commands


## Links
This is a way to store commonly used links in a single place. You can add, remove, and list links and commands.

### Add a link
```bash
show links add-item       
```
**Input**
```bash
URL: https://www.google.com
Name: Google
```

### Show all links
```bash
show links
```
Output
```bash
[0]Default group
  - [0]Stack overflow https://stackoverflow.com
  - [1]Google https://google.com
```

### Remove a link
```bash
-- removes the link at index 0.0 (groupId.itemIndex)
show links remove-item 0.0
```

### Open a link
```bash
-- opens the link at index 0.0 (groupId.itemIndex)
show links open 0.0
```

## Commands
This is a way to store commonly used commands in a single place. You can add, remove, and list commands.

### Add a command
```bash
show cli add-item       
```
**Input:**
```bash
Name: List in s3 
Code: aws s3 ls s3://my-bucket-name 
Item added to group [0]
```

### Show all commands
```bash
show cli
```
**Output:**
```bash
[0]Default Group
  [0][List Directory] ->  ls | grep golang 
  [1][List in s3] ->  aws s3 ls s3://my-bucket-name 
```

### Remove a command
```bash
-- removes the command at index 0.0 (groupId.itemIndex)
show cli remove-item 0.0
```

### Copy a command to clipboard
```bash
-- copies the command at index 0.0 (groupId.itemIndex) to clipboard
show cli copy 0.0
```

## Groups
The power of these tools is that you can have many groups and groups of groups and groups of groups of groups. You can add, remove, and list groups.
These groups are specific to links and commands.

### Add a group
```bash
-- shorthand: (show l ag)
show groups add-group       
```

**Input:**
```bash
Enter group name: Work
```

### Show all links
This will show all groups
```bash
-- shorthand: (show l)
show links
```

**Output:**
```bash
[0]Default group
  - [0]Stack overflow https://stackoverflow.com
  - [1]Google https://google.com

[1]Work
```

### Add a sub-group
If you do not specific a "--group-id" flag, the group will be added to the default group. In the following example I will add a sub group
to group 1.
```bash
-- shorthand: (show l asg --group-id=1)
show groups add-sub-group --group-id=1       
```

**Input:**
```bash
Enter group name: Work Sub group 
```

**Show all:**
```bash
-- shorthand: (show l)
show links
```

**Output:**
```bash
[0]Default group
  - [0]Stack overflow https://stackoverflow.com
  - [1]Google https://google.com

[1]Work
  - [0]Work Sub group
```

### For fun - add sub group of sub group and add a link
```bash
-- add sub group to sub-group
show l asg --group-id=1.0

-- add link to sub-sub-group
-- group-id=1.0.0 -> group 1, sub-group 0, sub-sub-group 0
show l ai --group-id=1.0.0

show l
output: 
[0]Default group
  - [0]Stack overflow https://stackoverflow.com
  - [1]Google https://google.com

[1]Work
  [0]Work Sub group
    [0]Work Sub-sub group
    - [0]Stack overflow https://stackoverflow.com

-- remove deep link
show l ri 1.0.0.0
Are you sure you want to remove item [0] from group [1.0.0]?
Enter 'y' to confirm or any other key to cancel
>y
Item removed
```