import 'package:flutter/material.dart';
import 'package:url_launcher/url_launcher.dart';

class MakerDrawer extends StatelessWidget {
  MakerDrawer({
    super.key,
    required this.labels,
    required this.icons,
    required this.changeIndex,
    this.child,
  });

  final double padding = 4;
  final drawerImage = "https://avatars.githubusercontent.com/u/174174755?v=4";
  final url = Uri.parse("https://github.com/segnities007");
  final double imageSize = 128;

  final List<String> labels;
  final List<IconData> icons;
  final void Function(int) changeIndex;
  final Widget? child;

  @override
  Widget build(context) {
    return Drawer(
        child: ListView(
            children: [
              DrawerHeader(
                decoration: const BoxDecoration(color: Colors.green),
                child: Center(
                  child: Container(
                    width: imageSize,
                    height: imageSize,
                    decoration: const BoxDecoration(
                      shape: BoxShape.circle,
                    ),
                    child: GestureDetector(
                      onTap: ()async{ await launchUrl(url); },
                      child: ClipOval(
                        child: Image.network(
                          drawerImage,
                          fit: BoxFit.cover,
                        ),
                      ),
                    ),
                  ),
                ),
              ),
              for(int i=0; i<labels.length; i++)
                MakerListTile(label: labels[i], icon: icons[i], changeIndex: changeIndex, index: i)
            ]
        )
    );
  }
}

class MakerListTile extends StatelessWidget {
  const MakerListTile({
    super.key,
    required this.label,
    required this.icon,
    required this.changeIndex,
    required this.index,
  });

  final int index;
  final String label;
  final IconData icon;
  final void Function(int) changeIndex;

  @override
  Widget build(context) {
    return ListTile(
        title: Padding(
            padding: const EdgeInsets.all(4),
            child: Row(
                children: [
                  const SizedBox(width: 10),
                  Icon(icon, size: 30),
                  const SizedBox(width: 10),
                  Text(label, style: const TextStyle(fontSize: 20)),
                ])
        ),
        onTap: () {
          changeIndex(index);
          Navigator.pop(context);
        });
  }
}
