import 'package:flutter/material.dart';

class MakerNavigationBar extends StatelessWidget {
  const MakerNavigationBar({
    super.key,
    required this.index,
    required this.changeIndex,
    required this.icons,
    required this.labels,
  });

  final int index;
  final void Function(int) changeIndex;
  final List<IconData> icons;
  final List<String> labels;

  @override
  Widget build(context) {
    return NavigationBar(
      backgroundColor: Colors.green,
      onDestinationSelected: changeIndex,
      selectedIndex: index,
      destinations: [
        for(int i=0; i<icons.length; i++)
          NavigationDestination(
            icon: Icon(icons[i]),
            label: labels[i],
        )
      ]
    );
  }
}
