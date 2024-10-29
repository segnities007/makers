import 'package:flutter/material.dart';

class MakerButton extends StatelessWidget {
  const MakerButton({
    super.key,
    required this.handler,
    required this.label,
    this.elevation = 8,
    this.minSize = 100,
  });

  final void Function() handler;
  final String label;
  final double elevation;
  final double minSize;

  @override
  Widget build(context) {
    return ElevatedButton(
      onPressed: handler,
      style: ElevatedButton.styleFrom(
        elevation: elevation,
        backgroundColor: Colors.green,
        minimumSize: Size(minSize, minSize*0.3),
      ),
      child: Text(label),
    );
  }
}
