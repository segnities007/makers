import 'package:flutter/material.dart';
import 'package:flutter_hooks/flutter_hooks.dart';

class MakerCard extends HookWidget {
  const MakerCard({
    super.key,
    required this.label,
    required this.description,
    this.padding = 16,
    this.radius = 16,
    this.iconSize = 32,
    this.isClickable = true,
    this.maxWidth = double.infinity,
  });

  final String label;
  final String description;
  final double padding;
  final double radius;
  final double iconSize;
  final bool isClickable;
  final double maxWidth;

  @override
  Widget build(BuildContext context) {
    return ConstrainedBox(
      constraints: BoxConstraints(
        maxWidth: maxWidth,
      ),
      child: Card(
        elevation: 4,
        margin: EdgeInsets.symmetric(vertical: padding / 2),
        shape: RoundedRectangleBorder(
          borderRadius: BorderRadius.circular(radius),
        ),
        color: Colors.white,
        child: InkWell(
          borderRadius: BorderRadius.circular(radius),
          onTap: isClickable
              ? () {
            //TODO:
          }
              : null,
          child: Padding(
            padding: EdgeInsets.all(padding),
            child: Row(
              children: [
                Icon(
                  Icons.group,
                  color: Colors.green,
                  size: iconSize,
                ),
                const SizedBox(width: 16),
                Expanded(
                  child: Column(
                    crossAxisAlignment: CrossAxisAlignment.start,
                    children: [
                      Text(
                        label,
                        style: const TextStyle(
                          fontSize: 32,
                          fontWeight: FontWeight.bold,
                        ),
                      ),
                      const SizedBox(height: 4),
                      Text(
                        description,
                        style: const TextStyle(
                          fontSize: 16,
                          color: Colors.grey,
                        ),
                      ),
                    ],
                  ),
                ),
              ],
            ),
          ),
        ),
      ),
    );
  }
}
