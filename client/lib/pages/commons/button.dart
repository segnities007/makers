import 'package:flutter/material.dart';

class MakerButton extends StatelessWidget {
  const MakerButton({
    super.key,
    required this.handler,
    required this.label,
    this.elevation = 8,
    this.size = 100,
    this.borderRadius = 8,
    this.padding = 16,
    this.height = 0.4,
  });

  final void Function() handler;
  final String label;
  final double elevation;
  final double size;
  final double borderRadius;
  final double padding;
  final double height;

  @override
  Widget build(BuildContext context) {
    return Container(
      padding: EdgeInsets.all(padding),
      width: size + padding * 2,
      height: size * height + padding * 2,
      child: ElevatedButton(
        onPressed: handler,
        style: ElevatedButton.styleFrom(
          elevation: elevation,
          backgroundColor: Colors.green,
          shape: RoundedRectangleBorder(
            borderRadius: BorderRadius.circular(borderRadius),
          ),
        ),
        child: Text(label, style: const TextStyle(color: Colors.white)),
      ),
    );
  }
}
