package playground.spring.sia.chapterthreejpa.tacocloud.controller;

import org.springframework.stereotype.Controller;
import org.springframework.web.bind.annotation.GetMapping;

@Controller
public class AboutController {

    @GetMapping("/about")
    public String home() {

        // returns the view name
        return "about";
    }
}
