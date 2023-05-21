import 'dart:convert';

import 'package:stylebooking_app/products.dart';
import 'package:flutter/material.dart';
import 'package:http/http.dart' as http;
import 'package:stylebooking_app/src/models/product.dart';

class ProductManager extends StatefulWidget {
  const ProductManager({super.key});

  @override
  State<StatefulWidget> createState() {
    return _ProductManagerState();
  }
}

class _ProductManagerState extends State<ProductManager> {
  var _products = List<ProductViewModel>.empty();

  @override
  void initState() {
    var products = fetchProducts();
    products.then((value) => setState(() {
      _products = value.products;
    }));  

    super.initState();
  }

  @override
  Widget build(BuildContext context) {
    return Products(_products);
  }

  Future<ProductListViewModel> fetchProducts() async {
    final response = await http
        .get(Uri.parse('http://localhost:8080/products'));

    if (response.statusCode == 200) {
      return ProductListViewModel.fromJson(jsonDecode(response.body));
    } else {
      // If the server did not return a 200 OK response,
      // then throw an exception.
      throw Exception('Failed to load products');
    }
  }
}