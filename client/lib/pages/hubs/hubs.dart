import 'package:flutter/material.dart';
import 'package:flutter_hooks/flutter_hooks.dart';

import '../commons/app_bar.dart';
import '../commons/navigation_bar.dart';
import './home/home.dart';
import './notification/notification.dart' as notify;
import './make/make.dart';
import './search/search.dart';
import './profile/profile.dart';
import '../commons/drawer.dart';

////

const icons = [
  Icons.home,
  Icons.search,
  Icons.add_circle,
  Icons.notifications,
  Icons.person,
];
const labels = ["Home", "Search", "Make", "Notify", "Profile"];
final hubs = [Home(), Search(), const Make(), const notify.Notification(), const Profile()];

final data = (
  labels: labels,
  icons: icons,
  hubs: hubs,
);
////

class Hubs extends HookWidget {
  const Hubs({super.key});

  @override
  Widget build(BuildContext context) {
    final index = useState(0);

    void Function(int) changeIndex() {
      return (int a) {
        index.value = a;
      };
    }

    return Scaffold(
      appBar: const MakerAppBar(title: "Makers"),
      drawer: MakerDrawer(
        labels: data.labels,
        icons: data.icons,
        changeIndex: changeIndex()
      ),
      body: data.hubs[index.value],
      bottomNavigationBar: MakerNavigationBar(
          labels: data.labels,
          icons: data.icons,
          index: index.value,
          changeIndex: changeIndex()
      )
    );
  }
}
