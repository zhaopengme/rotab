-- +migrate Up

PRAGMA foreign_keys = false;

-- ----------------------------
-- Table structure for rotab_archives
-- ----------------------------
DROP TABLE IF EXISTS "rotab_archives";
CREATE TABLE "rotab_archives"
(
    "id"            integer primary key autoincrement,
    "created_at"    datetime,
    "updated_at"    datetime,
    "deleted_at"    datetime,
    "year"          varchar(4),
    "month"         varchar(2),
    "article_count" integer,
    "blog_id"       bigint
);

-- ----------------------------
-- Table structure for rotab_articles
-- ----------------------------
DROP TABLE IF EXISTS "rotab_articles";
CREATE TABLE "rotab_articles"
(
    "id"            integer primary key autoincrement,
    "created_at"    datetime,
    "updated_at"    datetime,
    "deleted_at"    datetime,
    "author_id"     bigint,
    "title"         varchar(128),
    "abstract"      mediumtext,
    "tags"          text,
    "content"       mediumtext,
    "path"          varchar(255),
    "status"        integer,
    "topped"        bool,
    "commentable"   bool,
    "view_count"    integer,
    "comment_count" integer,
    "ip"            varchar(128),
    "user_agent"    varchar(255),
    "pushed_at"     datetime,
    "blog_id"       bigint
);

-- ----------------------------
-- Table structure for rotab_categories
-- ----------------------------
DROP TABLE IF EXISTS "rotab_categories";
CREATE TABLE "rotab_categories"
(
    "id"               integer primary key autoincrement,
    "created_at"       datetime,
    "updated_at"       datetime,
    "deleted_at"       datetime,
    "title"            varchar(128),
    "path"             varchar(255),
    "description"      varchar(255),
    "meta_keywords"    varchar(255),
    "meta_description" text,
    "tags"             text,
    "number"           integer,
    "blog_id"          bigint
);

-- ----------------------------
-- Table structure for rotab_comments
-- ----------------------------
DROP TABLE IF EXISTS "rotab_comments";
CREATE TABLE "rotab_comments"
(
    "id"                integer primary key autoincrement,
    "created_at"        datetime,
    "updated_at"        datetime,
    "deleted_at"        datetime,
    "article_id"        bigint,
    "author_id"         bigint,
    "content"           text,
    "parent_comment_id" bigint,
    "ip"                varchar(128),
    "user_agent"        varchar(255),
    "pushed_at"         datetime,
    "author_name"       varchar(32),
    "author_avatar_url" varchar(255),
    "author_url"        varchar(255),
    "blog_id"           bigint
);

-- ----------------------------
-- Table structure for rotab_correlations
-- ----------------------------
DROP TABLE IF EXISTS "rotab_correlations";
CREATE TABLE "rotab_correlations"
(
    "id"         integer primary key autoincrement,
    "created_at" datetime,
    "updated_at" datetime,
    "deleted_at" datetime,
    "id1"        bigint,
    "id2"        bigint,
    "str1"       varchar(255),
    "str2"       varchar(255),
    "str3"       varchar(255),
    "str4"       varchar(255),
    "int1"       integer,
    "int2"       integer,
    "int3"       integer,
    "int4"       integer,
    "type"       integer,
    "blog_id"    bigint
);

-- ----------------------------
-- Table structure for rotab_navigations
-- ----------------------------
DROP TABLE IF EXISTS "rotab_navigations";
CREATE TABLE "rotab_navigations"
(
    "id"          integer primary key autoincrement,
    "created_at"  datetime,
    "updated_at"  datetime,
    "deleted_at"  datetime,
    "title"       varchar(128),
    "url"         varchar(255),
    "icon_url"    varchar(255),
    "open_method" varchar(32),
    "number"      integer,
    "blog_id"     bigint
);

-- ----------------------------
-- Table structure for rotab_settings
-- ----------------------------
DROP TABLE IF EXISTS "rotab_settings";
CREATE TABLE "rotab_settings"
(
    "id"         integer primary key autoincrement,
    "created_at" datetime,
    "updated_at" datetime,
    "deleted_at" datetime,
    "category"   varchar(32),
    "name"       varchar(64),
    "value"      text,
    "blog_id"    bigint
);

-- ----------------------------
-- Table structure for rotab_tags
-- ----------------------------
DROP TABLE IF EXISTS "rotab_tags";
CREATE TABLE "rotab_tags"
(
    "id"            integer primary key autoincrement,
    "created_at"    datetime,
    "updated_at"    datetime,
    "deleted_at"    datetime,
    "title"         varchar(128),
    "article_count" integer,
    "blog_id"       bigint
);

-- ----------------------------
-- Table structure for rotab_users
-- ----------------------------
DROP TABLE IF EXISTS "rotab_users";
CREATE TABLE "rotab_users"
(
    "id"         integer primary key autoincrement,
    "created_at" datetime,
    "updated_at" datetime,
    "deleted_at" datetime,
    "username"   varchar(32),
    "nickname"   varchar(32),
    "password"   varchar(32),
    "email"      varchar(32),
    "salt"       varchar(32)
);

-- ----------------------------
-- Indexes structure for table rotab_archives
-- ----------------------------
CREATE INDEX "idx_rotab_archives_blog_id"
    ON "rotab_archives" (
                         "blog_id" ASC
        );
CREATE INDEX "idx_rotab_archives_deleted_at"
    ON "rotab_archives" (
                         "deleted_at" ASC
        );

-- ----------------------------
-- Auto increment value for rotab_articles
-- ----------------------------

-- ----------------------------
-- Indexes structure for table rotab_articles
-- ----------------------------
CREATE INDEX "idx_rotab_articles_blog_id"
    ON "rotab_articles" (
                         "blog_id" ASC
        );
CREATE INDEX "idx_rotab_articles_created_at"
    ON "rotab_articles" (
                         "created_at" ASC
        );
CREATE INDEX "idx_rotab_articles_deleted_at"
    ON "rotab_articles" (
                         "deleted_at" ASC
        );
CREATE INDEX "idx_rotab_articles_path"
    ON "rotab_articles" (
                         "path" ASC
        );
CREATE INDEX "idx_rotab_articles_status"
    ON "rotab_articles" (
                         "status" ASC
        );

-- ----------------------------
-- Indexes structure for table rotab_categories
-- ----------------------------
CREATE INDEX "idx_rotab_categories_blog_id"
    ON "rotab_categories" (
                           "blog_id" ASC
        );
CREATE INDEX "idx_rotab_categories_deleted_at"
    ON "rotab_categories" (
                           "deleted_at" ASC
        );

-- ----------------------------
-- Auto increment value for rotab_comments
-- ----------------------------

-- ----------------------------
-- Indexes structure for table rotab_comments
-- ----------------------------
CREATE INDEX "idx_rotab_comments_blog_id"
    ON "rotab_comments" (
                         "blog_id" ASC
        );
CREATE INDEX "idx_rotab_comments_deleted_at"
    ON "rotab_comments" (
                         "deleted_at" ASC
        );

-- ----------------------------
-- Indexes structure for table rotab_correlations
-- ----------------------------
CREATE INDEX "idx_rotab_correlations_blog_id"
    ON "rotab_correlations" (
                             "blog_id" ASC
        );
CREATE INDEX "idx_rotab_correlations_deleted_at"
    ON "rotab_correlations" (
                             "deleted_at" ASC
        );

-- ----------------------------
-- Indexes structure for table rotab_navigations
-- ----------------------------
CREATE INDEX "idx_rotab_navigations_blog_id"
    ON "rotab_navigations" (
                            "blog_id" ASC
        );
CREATE INDEX "idx_rotab_navigations_deleted_at"
    ON "rotab_navigations" (
                            "deleted_at" ASC
        );
-- ----------------------------
-- Indexes structure for table rotab_settings
-- ----------------------------
CREATE INDEX "idx_rotab_settings_blog_id"
    ON "rotab_settings" (
                         "blog_id" ASC
        );
CREATE INDEX "idx_rotab_settings_category"
    ON "rotab_settings" (
                         "category" ASC
        );
CREATE INDEX "idx_rotab_settings_deleted_at"
    ON "rotab_settings" (
                         "deleted_at" ASC
        );
CREATE INDEX "idx_rotab_settings_name"
    ON "rotab_settings" (
                         "name" ASC
        );


-- ----------------------------
-- Indexes structure for table rotab_tags
-- ----------------------------
CREATE INDEX "idx_rotab_tags_blog_id"
    ON "rotab_tags" (
                     "blog_id" ASC
        );
CREATE INDEX "idx_rotab_tags_deleted_at"
    ON "rotab_tags" (
                     "deleted_at" ASC
        );


-- ----------------------------
-- Indexes structure for table rotab_users
-- ----------------------------
CREATE INDEX "idx_rotab_users_deleted_at"
    ON "rotab_users" (
                      "deleted_at" ASC
        );

PRAGMA foreign_keys = true;

-- +migrate Down
