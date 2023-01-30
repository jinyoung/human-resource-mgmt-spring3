package human.resource.mgmt.common;

import human.resource.mgmt.EmployeeApplication;
import io.cucumber.spring.CucumberContextConfiguration;
import org.springframework.boot.test.context.SpringBootTest;

@CucumberContextConfiguration
@SpringBootTest(classes = { EmployeeApplication.class })
public class CucumberSpingConfiguration {}
