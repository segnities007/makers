import 'package:flutter/material.dart';
import 'package:flutter_hooks/flutter_hooks.dart';

import 'sign_in.dart';
import 'sign_up.dart';
import '../commons/app_bar.dart';
import '../commons/navigation_bar.dart';
import '../commons/drawer.dart';

////

final logins = [
  SignIn(),
  SignUp(),
];
const labels = [
  "sign in",
  "sign up",
];
const icons = [Icons.login, Icons.add];
final data = (logins: logins, labels: labels, icons: icons);

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
        body: Column(
          mainAxisAlignment: MainAxisAlignment.center,

          children: [
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
