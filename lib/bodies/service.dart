class Service {
  final Duration duration;
  final DateTime starttime;
  final String location;
  final String description;

  Service(this.duration, this.starttime, this.location, this.description);

  Map<String, dynamic> toJson() {
    return {
      "Duration": duration,
      "StartTime": starttime,
      "Location": location,
      "Description": description,
    };
  }
}