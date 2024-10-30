import 'package:flutter/material.dart';

class MakerButton extends StatelessWidget {
  const MakerButton({
    super.key,
    required this.handler,
    required this.label,
    this.elevation = 8,
    this.size = 100,
  });

  final void Function() handler;
  final String label;
  final double elevation;
  final double size;

  @override
  Widget build(BuildContext context) {
    return SizedBox(
      width: size,
      height: size * 0.3,
      child: ElevatedButton(
        onPressed: handler,
        style: ElevatedButton.styleFrom(
          elevation: elevation,
          backgroundColor: Colors.green,
          shape: RoundedRectangleBorder(
            borderRadius: BorderRadius.circular(4),
          ),
        ),
        child: Text(label, style: const TextStyle(color: Colors.white)),
      ),
    );
  }
}
