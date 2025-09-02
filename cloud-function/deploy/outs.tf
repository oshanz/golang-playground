output "function_uri" {
  value = google_cloudfunctions2_function.hello-function.service_config[0].uri
}