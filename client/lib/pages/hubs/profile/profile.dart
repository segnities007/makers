import 'package:flutter/material.dart';

import '../../commons/card.dart';

class Profile extends StatelessWidget {
  Profile({super.key});

  final double padding = 16;
  final List<IconData> icons = [
    Icons.edit,
    Icons.settings,
    Icons.logout,
  ];
  final List<String> labels = ["Setting", "Change Profile", "Log out"];

  @override
  Widget build(BuildContext context) {
    return Center(
        child: Container(
      padding: EdgeInsets.all(padding),
      child: Column(
        mainAxisAlignment: MainAxisAlignment.start,
        crossAxisAlignment: CrossAxisAlignment.center,
        children: [
          const MakerCard(
              iconSize: 128,
              label: "name",
              description: "description"),
          Column(mainAxisAlignment: MainAxisAlignment.center, children: [
            for (int i = 0; i < labels.length; i++)
              ProfileTile(
                  icon: icons[i],
                  label: labels[i],
                  handler: () {
                    //TODO
                  }),
          ])
        ],
      ),
    ));
  }
}

class ProfileTile extends StatelessWidget {
  const ProfileTile({
    super.key,
    required this.icon,
    required this.label,
    required this.handler,
    this.tileWidth = double.infinity,
    this.iconSize = 36,
    this.fontSize = 28,
    this.padding = 16,
  });

  final double padding;
  final IconData icon;
  final String label;
  final void Function() handler;
  final double tileWidth;
  final double iconSize;
  final double fontSize;

  @override
  Widget build(BuildContext context) {
    return InkWell(
      onTap: handler,
        child: Container(
          padding: EdgeInsets.all(padding),
          width: tileWidth,
          alignment: Alignment.center,
          child: Row(
            mainAxisAlignment: MainAxisAlignment.start,
            children: [
              Icon(icon, size: iconSize, color: Colors.grey[600]),
              const SizedBox(width: 8),
              Text(
                label,
                style: TextStyle(fontSize: fontSize, color: Colors.grey[600]),
              ),
            ],
          ),
      ),
    );
  }
}
