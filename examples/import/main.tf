terraform {
  required_providers {
    fazegallery = {
      version = "0.1"
      source  = "faze.com/gallery/faze-gallery"
    }
  }
}

resource "fazegallery_visualisation" "sample" {}

output "sample_vis" {
  value = fazegallery_visualisation.sample
}

