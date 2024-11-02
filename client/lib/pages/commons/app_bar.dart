import 'package:flutter/material.dart';

const double elevation = 0;

class MakerAppBar extends StatelessWidget implements PreferredSizeWidget{
  const MakerAppBar({
    super.key,
    required this.title
  });

  final String title;

  @override
  Widget build(context) {
    return AppBar(
      title: Center(
        child: Text(title)
      ),
      elevation: elevation,
      backgroundColor: Colors.green,
    );
  }

  @override
  Size get preferredSize => const Size.fromHeight(kToolbarHeight);
}
