package human.resource.mgmt.common;

import human.resource.mgmt.VacationApplication;
import io.cucumber.spring.CucumberContextConfiguration;
import org.springframework.boot.test.context.SpringBootTest;

@CucumberContextConfiguration
@SpringBootTest(classes = { VacationApplication.class })
public class CucumberSpingConfiguration {}
