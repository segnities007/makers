import 'package:flutter/material.dart';
import 'package:flutter_hooks/flutter_hooks.dart';

const double radius = 16;
const double padding = 16;

class MakerCard extends HookWidget {
  const MakerCard({super.key, required this.index});

  final int index;

  @override
  Widget build(BuildContext context) {
    return Card(
      margin: const EdgeInsets.symmetric(vertical: padding/2),
      shape: RoundedRectangleBorder(
        borderRadius: BorderRadius.circular(radius),
      ),
      color: Colors.white,
      child: InkWell(
        borderRadius: BorderRadius.circular(radius),
        onTap: () {
          //TODO
        },
        child: const Padding(
          padding: EdgeInsets.all(padding),
          child: Row(
            children: [
              Icon(Icons.group, color: Colors.green),
              SizedBox(width: 16),
              Expanded(
                child: Column(
                  crossAxisAlignment: CrossAxisAlignment.start,
                  children: [
                    Text(
                      "Title",
                      style: TextStyle(
                        fontSize: 32,
                        fontWeight: FontWeight.bold,
                      ),
                    ),
                    SizedBox(height: 4),
                    Text(
                      "description",
                      style: TextStyle(
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
    );
  }
}
