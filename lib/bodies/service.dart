class Service {
  final Duration duration;
  final String location;
  final String description;

  Service(this.duration, this.location, this.description);

  Map<String, dynamic> toJson() {
    return {
      "Duration": duration,
      "Location": location,
      "Description": description,
    };
  }
}
