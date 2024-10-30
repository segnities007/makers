import 'package:flutter/material.dart';
import 'package:flutter_hooks/flutter_hooks.dart';
import 'package:go_router/go_router.dart';

import 'sign_in.dart';
import 'sign_up.dart';
import '../commons/app_bar.dart';
import '../commons/navigation_bar.dart';
import '../commons/button.dart';
import '../commons/drawer.dart';

////

const logins = [
  SignIn(),
  SignUp(),
];
const labels = [
  "sign in",
  "sign up",
];
const icons = [Icons.login, Icons.add];
const data = (logins: logins, labels: labels, icons: icons);

////

class Logins extends HookWidget {
  const Logins({super.key});

  @override
  Widget build(BuildContext context) {
    final index = useState(0);
    void changeIndex(int a) {
        index.value = a;
    }

    return Scaffold(
        appBar: const MakerAppBar(
          title: "Login",
        ),
        drawer: MakerDrawer(
            labels: data.labels, icons: data.icons, changeIndex: changeIndex),
        body: Column(children: [
          Center(
            child: MakerButton(
                handler: () {
                  context.go("/hubs");
                },
                label: "go hubs",
                size: 200),
          ),
          data.logins[index.value]
        ]),
        bottomNavigationBar: MakerNavigationBar(
          index: index.value,
          changeIndex: changeIndex,
          icons: data.icons,
          labels: data.labels,
        ));
  }
}
