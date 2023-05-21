import 'package:stylebooking_app/products.dart';
import 'package:flutter/material.dart';

class ProductManager extends StatefulWidget {
  const ProductManager({super.key});

  @override
  State<StatefulWidget> createState() {
    return _ProductManagerState();
  }
}

class _ProductManagerState extends State<ProductManager> {
  final List<String> _products = [
    'Cut',
    'Nails',
    'Wax',
  ];

  @override
  void initState() {
    super.initState();
  }

  @override
  Widget build(BuildContext context) {
    return Products(_products);
  }
}