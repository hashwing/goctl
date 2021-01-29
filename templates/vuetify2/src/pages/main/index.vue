//src/pages/main/index.vue
<template>
  <v-app id="app">
    <v-navigation-drawer v-model="drawer" clipped fixed app>
      <v-list dense>
        <v-list-group
          v-for="item in items"
          :key="item.title"
          v-model="item.active"
          :prepend-icon="item.action"
          no-action
        >
          <template v-slot:activator>
            <v-list-item-content>
              <v-list-item-title v-text="item.title"></v-list-item-title>
            </v-list-item-content>
          </template>

          <v-list-item
            v-for="subItem in item.items"
            :key="subItem.title"
            :to="subItem.path"
          >
            <v-list-item-content>
              <v-list-item-title v-text="subItem.title"></v-list-item-title>
            </v-list-item-content>
          </v-list-item>
        </v-list-group>
      </v-list>
    </v-navigation-drawer>
    <v-app-bar app fixed dark clipped-left color="indigo">
      <v-app-bar-nav-icon @click.stop="drawer = !drawer"></v-app-bar-nav-icon>
      <v-toolbar-title>{{ .App }}</v-toolbar-title>
      <v-spacer></v-spacer>
      <v-menu min-width="200" offset-y>
        <template v-slot:activator="{ on, attrs }">
          <v-avatar size="36" color="blue darken-1" dark v-bind="attrs" v-on="on">
            <v-icon dark>mdi-account-circle</v-icon>
          </v-avatar>
        </template>
        <v-list dense>
          <v-list-item link>
            <v-list-item-content>
              <v-list-item-title>游客</v-list-item-title>
              <v-list-item-subtitle>anymous@example.com</v-list-item-subtitle>
            </v-list-item-content>
          </v-list-item>
          <v-list-item>
            <v-divider></v-divider>
          </v-list-item>
          <v-list-item link>
            <v-list-item-icon>
              <v-icon>mdi-logout</v-icon>
            </v-list-item-icon>
            <v-list-item-title>退出</v-list-item-title>
          </v-list-item>
        </v-list>
      </v-menu>
    </v-app-bar>
    <v-content>
      <v-container fluid>
        <router-view></router-view>
      </v-container>
    </v-content>
  </v-app>
</template>

<script>
export default {
  data: () => ({
    drawer: true,
    items: [
      {
        action: "mdi-view-dashboard",
        active: true,
        title: "概览",
        items: [{ title: "系统概览", path: "/" }],
      },
      {
        action: "mdi-settings",
        title: "设置",
        items: [
          { title: "系统设置", path: "/setting" },
        ],
      },
    ],
    itemIndex: 0,
  }),
  components: {},
};
</script>

<style >
html {
  overflow-y: auto !important;
}
</style>