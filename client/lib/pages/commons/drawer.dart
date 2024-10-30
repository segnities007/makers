import 'package:flutter/material.dart';

////

const double padding = 4;

////

class MakerDrawer extends StatelessWidget {
  const MakerDrawer({
    super.key,
    required this.labels,
    required this.icons,
    required this.changeIndex,
  });

  final List<String> labels;
  final List<IconData> icons;
  final void Function(int) changeIndex;

  @override
  Widget build(context) {
    return Drawer(
        child: ListView(
          children: [
      const DrawerHeader(
          decoration: BoxDecoration(color: Colors.green),
          child: Text("Header")),
        for(int i=0; i<labels.length; i++)
          MakerListTile(label: labels[i], icon: icons[i], changeIndex: changeIndex, index: i)
    ]));
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
          // mainAxisAlignment: MainAxisAlignment.center,
          children: [
            const SizedBox(width: 10),
            Icon(icon, size: 30),
            const SizedBox(width: 10),
            Text(label + index.toString(), style: const TextStyle(fontSize: 20)),
        ])
      ),
        onTap: () {
          changeIndex(index);
          Navigator.pop(context);
        });
  }
}
