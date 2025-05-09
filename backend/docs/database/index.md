# Database Schema

Structure and relationships of the SQLite database tables used in the MinBan system.

![Database Schema Diagram](../img/database/database.svg)

## Tables

### **User Table**
Stores information about users.

| Column   | Type        | Constraints       |
|----------|-------------|-------------------|
| `id`     | char(36)    | Primary Key       |
| `name`   | varchar(40) | NOT NULL          |
| `password` | char(64)  | NOT NULL          |

---

### **Card Table**
Stores information about cards.

| Column      | Type        | Constraints       |
|-------------|-------------|-------------------|
| `id`        | char(36)    | Primary Key       |
| `title`     | varchar(60) | NOT NULL          |
| `description` | text      | NOT NULL          |
| `position`  | int         | NOT NULL          |
| `state`     | int         | Foreign Key       |
| `user_id`   | char(36)    | Foreign Key       |

---

### **Tag Table**
Stores information about tags.

| Column   | Type        | Constraints       |
|----------|-------------|-------------------|
| `name`   | varchar(20) | Primary Key       |
| `color`  | varchar(6)  | NOT NULL          |

---

### **Card_Tags Table**
Associates cards with tags. This table is used to manage the many-to-many relationship between cards and tags.

| Column    | Type        | Constraints       |
|-----------|-------------|-------------------|
| `card_id` | varchar(36) | Foreign Key       |
| `tag`     | varchar(20) | Foreign Key       |

---

### **State Table**
Stores information about states (columns) for organizing cards.

| Column     | Type        | Constraints       |
|------------|-------------|-------------------|
| `id`       | int         | Primary Key       |
| `name`     | varchar(20) | NOT NULL          |
| `position` | int         | NOT NULL          |
| `color`    | char(6)     | NOT NULL          |

