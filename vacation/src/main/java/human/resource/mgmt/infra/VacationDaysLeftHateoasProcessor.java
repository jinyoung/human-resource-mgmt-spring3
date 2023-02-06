package human.resource.mgmt.infra;

import human.resource.mgmt.domain.*;
import org.springframework.hateoas.EntityModel;
import org.springframework.hateoas.Link;
import org.springframework.hateoas.server.RepresentationModelProcessor;
import org.springframework.stereotype.Component;

@Component
public class VacationDaysLeftHateoasProcessor
    implements RepresentationModelProcessor<EntityModel<VacationDaysLeft>> {

    @Override
    public EntityModel<VacationDaysLeft> process(
        EntityModel<VacationDaysLeft> model
    ) {
        model.add(
            Link
                .of(model.getRequiredLink("self").getHref() + "/add")
                .withRel("add")
        );
        model.add(
            Link
                .of(model.getRequiredLink("self").getHref() + "/use")
                .withRel("use")
        );

        return model;
    }
}
