package human.resource.mgmt.infra;

import human.resource.mgmt.domain.*;
import org.springframework.hateoas.EntityModel;
import org.springframework.hateoas.Link;
import org.springframework.hateoas.server.RepresentationModelProcessor;
import org.springframework.stereotype.Component;

@Component
public class EmployeeHateoasProcessor
    implements RepresentationModelProcessor<EntityModel<Employee>> {

    @Override
    public EntityModel<Employee> process(EntityModel<Employee> model) {
        model.add(
            Link
                .of(model.getRequiredLink("self").getHref() + "/resign")
                .withRel("resign")
        );

        return model;
    }
}
