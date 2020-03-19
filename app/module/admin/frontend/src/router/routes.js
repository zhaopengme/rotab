import Login from "../views/Login.vue";
import Main from "../components/Main.vue";
import Container from "../components/Container.vue";

const routes = [
  {
    path: "/",
    name: "login",
    component: Login
  },
  {
    path: "/app",
    name: "app",
    component: Main,
    children: [
      {
        path: "/home",
        name: "控制台",
        component: () => import("../views/Home.vue")
      },
      {
        path: "/post",
        name: "新建文章",
        redirect: "/post/edit"
      },
      {
        path: "/post",
        name: "文章",
        component: Container,
        children: [
          {
            path: "/post/edit",
            name: "新建文章",
            component: () => import("../views/post/edit.vue")
          },
          {
            path: "/post/list",
            name: "文章列表",
            component: () => import("../views/post/list.vue")
          },
          {
            path: "/post/comments",
            name: "评论",
            component: () => import("../views/post/comments.vue")
          },
          {
            path: "/post/tags",
            name: "标签",
            component: () => import("../views/post/tags.vue")
          },
          {
            path: "/post/categories",
            name: "分类",
            component: () => import("../views/post/categories.vue")
          },
          {
            path: "/post/navigations",
            name: "导航",
            component: () => import("../views/post/navigations.vue")
          }
        ]
      },
      {
        path: "/media",
        name: "媒体",
        component: Container,
        children: [
          {
            path: "/media/list",
            name: "媒体库",
            component: () => import("../views/media/list.vue")
          },
          {
            path: "/media/new",
            name: "添加",
            component: () => import("../views/media/new.vue")
          }
        ]
      },
      {
        path: "/themes",
        name: "主题",
        component: Container,
        children: [
          {
            path: "/themes/list",
            name: "主题库",
            component: () => import("../views/themes/list.vue")
          }
        ]
      },
      {
        path: "/users",
        name: "用户",
        component: Container,
        children: [
          {
            path: "/users/list",
            name: "个人资料",
            component: () => import("../views/users/profile.vue")
          },
          {
            path: "/users/password",
            name: "修改密码",
            component: () => import("../views/users/password.vue")
          }
        ]
      },
      {
        path: "/options",
        name: "设置",
        component: Container,
        children: [
          {
            path: "/options/general",
            name: "常规",
            component: () => import("../views/options/general.vue")
          },
          {
            path: "/options/reading",
            name: "阅读",
            component: () => import("../views/options/reading.vue")
          },
          {
            path: "/options/writing",
            name: "撰写",
            component: () => import("../views/options/writing.vue")
          },
          {
            path: "/options/discussion",
            name: "讨论",
            component: () => import("../views/options/discussion.vue")
          }
        ]
      }
    ]
  }
];

export default routes;
